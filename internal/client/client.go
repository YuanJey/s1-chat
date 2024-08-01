package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type TCPClient struct {
	Address string
	conn    net.Conn
}

func NewTCPClient(address string) *TCPClient {
	return &TCPClient{Address: address}
}

func (c *TCPClient) ConnectToServer() {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	//defer conn.Close()
	c.conn = conn
	fmt.Println("Connected to server. Type messages to send.")

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			text := scanner.Text()
			_, err := conn.Write([]byte(text + "\n"))
			if err != nil {
				fmt.Println("Error sending message:", err)
				break
			}

			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading response:", err)
				break
			}
			fmt.Print("Server response: ", response)
		}
	}()

}
func (c *TCPClient) SendMsg(text string) {
	_, err := c.conn.Write([]byte(text + "\n"))
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
