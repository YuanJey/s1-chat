package cluster

import (
	"bufio"
	"fmt"
	"net"
	"s1-chat/pkg/consts"
)

type Cluster struct {
	registerServer *RegisterServer
}

func (c *Cluster) StartClusterServer() {
	listener, err := net.Listen("tcp", ":"+consts.ClusterServerPort)
	if err != nil {
		fmt.Printf("Error starting TCP server: %s\n", err)
		return
	}
	defer listener.Close()
	fmt.Printf("TCP server started on port %s\n", consts.ClusterServerPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		go c.handleConnection(conn)
	}
}
func (c *Cluster) handleConnection(conn net.Conn) {
	fmt.Printf("Connection accepted from %s\n", conn.RemoteAddr().String())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Bytes()
		fmt.Printf("Received: %s\n", message)
		c.Work(conn, message)
		_, err := conn.Write(message)
		if err != nil {
			fmt.Printf("Error sending response: %s\n", err)
			break
		}
	}
	conn.Close()
}
func (c *Cluster) Work(conn net.Conn, msg []byte) {
	//TODO
}
