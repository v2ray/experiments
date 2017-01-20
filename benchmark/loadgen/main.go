package main

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

var (
	fPort        = flag.Int("port", 10000, "Proxy port for sending traffic to.")
	fRemotePort  = flag.Int("remoteport", 10001, "Remote port.")
	fType        = flag.String("type", "direct", "Proxy type of the target, either 'direct' or 'socks'.")
	fAmount      = flag.Int("amount", 1, "Amount of traffic to send, in GB.")
	fConcurrency = flag.Int("concurrency", 1, "Number of concurrect connections for benchmark.")
)

func makeConnection() (net.Conn, error) {
	switch *fType {
	case "direct":
		return net.DialTCP("tcp4", nil, &net.TCPAddr{
			IP:   net.IP([]byte{127, 0, 0, 1}),
			Port: *fPort,
		})
	case "socks":
		dialer, err := proxy.SOCKS5("tcp4", fmt.Sprintf(":%d", *fPort), nil, proxy.Direct)
		if err != nil {
			return nil, err
		}
		return dialer.Dial("tcp4", fmt.Sprintf(":%d", *fRemotePort))
	default:
		return nil, errors.New("Unknown proxy type: " + *fType)
	}
}

func main() {
	flag.Parse()

	const BufSize = 128 * 1024
	var wg sync.WaitGroup

	startTime := time.Now().Unix()
	for i := 0; i < *fConcurrency; i++ {
		wg.Add(1)
		go func() {
			buf := make([]byte, BufSize)
			rand.Read(buf)
			conn, err := makeConnection()
			if err != nil {
				panic(err)
			}
			var connWg sync.WaitGroup
			connWg.Add(2)
			go func() {
				totalBytes := int64(*fAmount) * 1024 * 1024 * 1024
				for totalBytes > 0 {
					_, err := conn.Write(buf)
					if err != nil {
						panic(err)
					}
					totalBytes -= BufSize
				}
				connWg.Done()
			}()
			go func() {
				totalBytes := int64(*fAmount) * 1024 * 1024 * 1024
				for {
					var count uint64
					if err := binary.Read(conn, binary.BigEndian, &count); err != nil {
						panic(err)
					}
					if count >= uint64(totalBytes) {
						break
					}
				}
				connWg.Done()
			}()
			connWg.Wait()
			conn.Close()
			wg.Done()
		}()
	}
	wg.Wait()

	endTime := time.Now().Unix()
	elapsed := endTime - startTime
	if elapsed == 0 {
		fmt.Println("Finished in 0 second. Too fast for benchmark.")
		return
	}
	dataAmount := uint64(*fConcurrency) * uint64(*fAmount)

	speed := dataAmount * 1024 / uint64(elapsed)
	fmt.Println("LoadGen:", dataAmount, "GB of data sent in", elapsed, "seconds, with speed", speed, "MB/s.")
}
