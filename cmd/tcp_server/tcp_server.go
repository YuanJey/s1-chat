package main

import (
	"fmt"
	"s1-chat/internal/handle"
	"s1-chat/internal/server"
	"s1-chat/pkg/structs"
)

type TestHandle struct {
}

func (t *TestHandle) Processing(msg *structs.Message) bool {
	fmt.Println("tcp server Handle msg ")
	fmt.Println(msg)
	return true
}
func main() {
	tcpServer := server.NewMessageServer("8000")
	manage := handle.NewManage()
	manage.AddBeforeHandle(0, &TestHandle{})
	manage.AddBeforeHandle(1, &TestHandle{})
	tcpServer.SetManage(manage)
	tcpServer.StartServer()
}
