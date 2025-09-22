package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":8081")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	fmt.Println("UDP server listening on port 8081...")
	buf := make([]byte, 1024)

	for {
		n, clientAddr, _ := conn.ReadFromUDP(buf)
		msg := string(buf[:n])
		fmt.Println("Received:", msg)

		// Send response
		conn.WriteToUDP([]byte("Hello "+msg), clientAddr)
	}
}
