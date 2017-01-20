package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
)

var (
	fPort = flag.Int("port", 10002, "Port to listen on.")
)

func receive(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 128*1024)
	var total uint64
	var c int
	for {
		n, err := conn.Read(buf)
		total += uint64(n)
		if err != nil {
			fmt.Println("Connection finishes with", total, "bytes:", err)
			return
		}
		c++
		if c == 16 {
			c = 0
			if err := binary.Write(conn, binary.BigEndian, total); err != nil {
				panic(err)
			}
		}
	}
}

func accept(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}
		go receive(conn)
	}
}

func main() {
	flag.Parse()
	listener, err := net.ListenTCP("tcp4", &net.TCPAddr{
		IP:   net.IP([]byte{127, 0, 0, 1}),
		Port: *fPort,
	})
	if err != nil {
		panic(err)
	}
	go accept(listener)
	c := make(chan bool)
	<-c
}
