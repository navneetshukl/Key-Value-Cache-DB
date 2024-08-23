// server.go
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type TCPserver struct {
}

func NewTCPServer() *TCPserver {
	return &TCPserver{}
}

func (s *TCPserver) handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from client:", err)
			}
			return
		}
		fmt.Printf("Received message: %s\n", string(buffer[:n]))

		_, err = conn.Write([]byte("Message received"))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
	}
}

func main() {
	// Listen on port 8080

	s := NewTCPServer()
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Println("Server listening on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go s.handleConnection(conn)
	}
}
