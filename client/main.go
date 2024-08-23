// client.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type TCPclient struct {
}

func NewTCPClient() *TCPclient {
	return &TCPclient{}
}

func (c *TCPclient) handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		log.Println("Enter the command : ")
		command, _ := reader.ReadString('\n')

		command = strings.TrimSpace(command)

		// send the command to the server
		_, err := conn.Write([]byte(command + "\n"))
		if err != nil {
			log.Println("Error sending the command :", err)
			return
		}

		// Read the server's response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		fmt.Printf("Response from server: %s\n", string(buffer[:n]))
	}

}

// func get() {
// 	kv := storage.NewKV()

// 	kv.Set("mykey", "myvalue")
// 	value, err := kv.Get("mykey")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Value:", value) // Should print "myvalue"
// 	}
// }

func main() {
	// Connect to server on port 8080
	c := NewTCPClient()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}

	c.handleClient(conn)

	// get()

}
