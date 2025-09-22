package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr, _ := net.ResolveUDPAddr("udp", "localhost:8081")
	conn, _ := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()

	conn.Write([]byte("World"))
	buf := make([]byte, 1024)
	n, _, _ := conn.ReadFromUDP(buf)
	fmt.Println("Server reply:", string(buf[:n]))
}
