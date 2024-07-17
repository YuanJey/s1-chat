package server

import (
	"bufio"
	"fmt"
	"net"
)

// StartTCPServer starts a TCP server on the specified port.
func StartTCPServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Error starting TCP server: %s\n", err)
		return
	}
	defer listener.Close()
	fmt.Printf("TCP server started on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

// handleConnection handles the incoming connections.
func handleConnection(conn net.Conn) {
	fmt.Printf("Connection accepted from %s\n", conn.RemoteAddr().String())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Received: %s\n", message)
		_, err := conn.Write([]byte("Received: " + message + "\n"))
		if err != nil {
			fmt.Printf("Error sending response: %s\n", err)
			break
		}
	}
	conn.Close()
}
