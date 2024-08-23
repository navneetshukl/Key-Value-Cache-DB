// server.go
package main

import (
	"fmt"
	"io"
	"key-value-db/storage"
	"log"
	"net"
	"os"
	"strings"
)

type TCPserver struct {
	KV storage.KeyValueDB
}

func NewTCPServer(kv storage.KeyValueDB) *TCPserver {
	return &TCPserver{
		KV: kv,
	}
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

		command := strings.TrimSpace(string(buffer[:n]))
		parts := strings.Fields(command)
		if len(parts) == 0 {
			continue
		}
		var response string
		switch parts[0] {
		case "SET":
			if len(parts) != 3 {
				response = "Usage: SET key value\n"
			} else {
				s.KV.Set(parts[1], parts[2])
				response = "OK\n"
			}
		case "GET":
			if len(parts) != 2 {
				response = "Usage: GET key\n"
			} else {
				value, err := s.KV.Get(parts[1])
				if err != nil {
					log.Println("Error in getting the key ", err)

					response = "(nil)\n"
				} else {
					response = value + "\n"
				}
			}

		default:
			response = "Unknown command\n"
		}
		// fmt.Printf("Received message: %s\n", string(buffer[:n]))

		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Error writing to client:", err)
			return
		}
	}
}

func main() {
	// Listen on port 8080

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
		storage := storage.NewKV()

		s := NewTCPServer(storage)
		log.Println("Address is ", s)
		go s.handleConnection(conn)
	}
}
