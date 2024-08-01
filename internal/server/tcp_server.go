package server

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"s1-chat/internal/handle"
	"s1-chat/pkg/structs"
	"s1-chat/pkg/utils"
)

type TCPServer struct {
	wsAddr  string
	port    string
	Manage  *handle.Manage
	connMap map[string]net.Conn
}

func NewMessageServer(port string) *TCPServer {
	return &TCPServer{port: port, connMap: make(map[string]net.Conn)}
}
func (s *TCPServer) SetManage(manage *handle.Manage) {
	s.Manage = manage
}
func (s *TCPServer) addConn(conn *net.Conn) {
}

// StartServer starts a TCP ws_server on the specified port.
func (s *TCPServer) StartServer() {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		fmt.Printf("Error starting TCP ws_server: %s\n", err)
		return
	}
	defer listener.Close()
	fmt.Printf("TCP ws_server started on port %s\n", s.port)

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
		go s.Work(message)
		_, err := conn.Write(message)
		if err != nil {
			fmt.Printf("Error sending response: %s\n", err)
			break
		}
	}
	conn.Close()
}
func (s *TCPServer) Work(msg []byte) {
	buff := bytes.NewBuffer(msg)
	message := structs.Message{}
	err := utils.JsonStringToStruct(buff.String(), &message)
	if err != nil {
		fmt.Printf(" Work err: %s\n", err)
		return
	}
	s.Manage.ProcessMessage(&message)
}
func (s *TCPServer) Send(toId string, msg structs.Msg) {
	if conn, ok := s.connMap[toId]; ok {
		_, err := conn.Write(msg.ToByte())
		if err != nil {
			fmt.Printf("Error sending response: %s\n", err)
		}
	}
}
