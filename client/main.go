// client.go
package main

import (
	"fmt"
	"net"
	"os"
)

type TCPclient struct {
}

func NewTCPClient() *TCPclient {
	return &TCPclient{}
}

func (c *TCPclient) handleClient(conn net.Conn) {
	defer conn.Close()

	message := "Hello, Server!"
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	fmt.Printf("Received from server: %s\n", string(buffer[:n]))
}

func main() {
	// Connect to server on port 8080
	c := NewTCPClient()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}

	c.handleClient(conn)

}
