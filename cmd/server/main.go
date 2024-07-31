package main

import "s1-chat/internal/server"

func main() {
	messageServer := server.NewMessageServer("8000")
	messageServer.StartTCPServer()
}
