package server

import (
	"bufio"
	"fmt"
	"net"
	"s1-chat/internal/handle"
	"s1-chat/pkg/structs"
	"s1-chat/pkg/utils"
)

type TCPServer struct {
	port   string
	manage *handle.Manage2
}

// StartTCPServer starts a TCP server on the specified port.
func (s *TCPServer) StartTCPServer() {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		fmt.Printf("Error starting TCP server: %s\n", err)
		return
	}
	defer listener.Close()
	fmt.Printf("TCP server started on port %s\n", s.port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

// handleConnection handles the incoming connections.
func (s *TCPServer) handleConnection(conn net.Conn) {
	fmt.Printf("Connection accepted from %s\n", conn.RemoteAddr().String())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Bytes()
		fmt.Printf("Received: %s\n", message)
		s.Work(conn, message)
		_, err := conn.Write(message)
		if err != nil {
			fmt.Printf("Error sending response: %s\n", err)
			break
		}
	}
	conn.Close()
}
func (s *TCPServer) Work(conn net.Conn, msg []byte) {
	message := structs.Message{}
	err := utils.JsonStringToStruct(string(msg), message)
	if err != nil {
		fmt.Printf(" Work err: %s\n", err)
		return
	}
	finish := s.manage.ProcessMessage(&message)
	if finish {
		finishMessage := structs.Message{SendFinishMessage: struct {
			Id string `json:"id"`
		}(struct{ Id string }{Id: message.Id})}
		s.Send(conn, finishMessage)
	}
}
func (s *TCPServer) Send(conn net.Conn, msg structs.Message) {
	_, err := conn.Write([]byte(utils.StructToJsonString(msg)))
	if err != nil {
		fmt.Printf("Error sending response: %s\n", err)
	}
}
