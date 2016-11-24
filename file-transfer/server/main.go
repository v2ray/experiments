package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"net"
)

var (
	port = flag.String("port", "", "Port to listen, e.g.: \":8080\"")
)

func saveFile(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024*1024)
	nBytes, err := io.ReadAtLeast(conn, buffer, 4)
	if err != nil {
		fmt.Println("Failed to read header:", err)
		return
	}
	filenameLen := binary.BigEndian.Uint32(buffer[:4])
	copy(buffer, buffer[4:nBytes])
	nBytes -= 4
	if uint32(nBytes) < filenameLen {
		_, err := io.ReadFull(conn, buffer[nBytes:filenameLen])
		if err != nil {
			fmt.Println("Failed to read header:", err)
		}
	}
	filename := string(buffer[:filenameLen])

	//f, err := os.Create(filename)
	//if err != nil {
	//	fmt.Println("Failed to create file:", err)
	//	conn.Write([]byte{0x01})
	//	return
	//}
	//defer f.Close()

	conn.Write([]byte{0x00})

	fmt.Println("Transfering file:", filename)
	crc := crc32.NewIEEE()
	totalBytesReceived := 0
	for {
		nBytes, err := conn.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Failed to read file content:", err)
			return
		}
		totalBytesReceived += nBytes
		crc.Write(buffer[:nBytes])
	}
	fmt.Println(filename, "received", totalBytesReceived, "bytes with crc:", crc.Sum32())
}

func main() {
	flag.Parse()

	listerner, err := net.Listen("tcp", *port)
	if err != nil {
		fmt.Println("Failed to listen on", *port, ":", err)
		return
	}

	for {
		conn, err := listerner.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			return
		}
		go saveFile(conn)
	}
}
