package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send message
	fmt.Fprintf(conn, "Hello Server\n")

	// Read reply
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Server reply:", string(buf[:n]))
}