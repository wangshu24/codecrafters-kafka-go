package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	
	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	buff := make([]byte, 1024)
	length, connErr := conn.Read(buff)
	message := string(buff)
	
	fmt.Println("String: ", message)
	
	if connErr != nil {
		fmt.Println("Error reading from connection: ", connErr.Error())
		os.Exit(1)
	}

	if length == 0 {
		fmt.Println("Empty message from connection")
		os.Exit(1)
	}

	msgSize := binary.BigEndian.Uint32(buff[:4])
	fmt.Println("message size: ", msgSize)

	apiKey := binary.BigEndian.Uint16(buff[4:6])
	fmt.Println("api key: ", apiKey)

	apiVer := binary.BigEndian.Uint16(buff[6:8])
	fmt.Println("api version: ", apiVer)

	corrId := binary.BigEndian.Uint32(buff[8:12])
	fmt.Println("correlation ID: ", corrId)
	resp := make([]byte, 8)
	copy(resp,[]byte{00,00,00,00})
	copy(resp[4:], buff[8:12])
	
	conn.Write([]byte(resp))
}
