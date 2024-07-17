package main

import "s1-chat/internal/server"

func main() {
	server.StartTCPServer("8000")
}
