package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

var (
	file   = flag.String("file", "", "File to transfer.")
	server = flag.String("server", "", "Server address and port, e.g.: \"8.8.8.8:8080\".")
)

func getHeader() []byte {
	filename := filepath.Base(*file)
	header := make([]byte, len(filename)+4)
	binary.BigEndian.PutUint32(header, uint32(len(filename)))
	copy(header[4:], []byte(filename))
	return header
}

func main() {
	flag.Parse()

	info, err := os.Stat(*file)
	if err != nil {
		fmt.Println("Failed to read file(", *file, "):", err)
		return
	}
	if info.IsDir() {
		fmt.Println("Unable to transfer directory:", *file)
		return
	}
	fileSize := info.Size()
	fmt.Println("File name:", *file, "Size:", fileSize/1024/1024, "MB")

	buffer := make([]byte, 1024*1024)

	fmt.Println("Initializing connection.")
	conn, err := net.Dial("tcp", *server)
	if err != nil {
		fmt.Println("Failed to connect to server(", *server, "):", err)
		return
	}
	defer conn.Close()

	if _, err := conn.Write(getHeader()); err != nil {
		fmt.Println("Failed to send header:", err)
		return
	}

	if _, err := conn.Read(buffer); err != nil || buffer[0] != 0 {
		fmt.Println("Server failed to respond.")
		return
	}

	startTime := time.Now().UTC()
	fmt.Println("Begin transfer file. ", startTime.Format(time.RFC1123Z))

	f, err := os.Open(*file)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer f.Close()

	totalBytesSent := int64(0)
	crc := crc32.NewIEEE()
	for {
		nBytes, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return
		}
		crc.Write(buffer[0:nBytes])

		bytesSent := 0
		for bytesSent < nBytes {
			n, err := conn.Write(buffer[bytesSent:nBytes])
			if err != nil {
				fmt.Println("Failed to send file:", err)
				return
			}
			bytesSent += n
		}
		totalBytesSent += int64(bytesSent)
	}

	if totalBytesSent != fileSize {
		fmt.Println("Failed to send whole file. Sent bytes", totalBytesSent, "file bytes", fileSize)
	}

	endTime := time.Now().UTC()
	timeElapsed := endTime.Sub(startTime) / time.Second
	fmt.Println(*file, "transfer finished at", endTime.Format(time.RFC1123Z), "crc", crc.Sum32(), "time elapsed: ", timeElapsed, "seconds.")
}
