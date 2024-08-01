package main

import (
	"s1-chat/internal/handle"
	"s1-chat/internal/server"
)

func main() {
	var ws server.WsServer
	manage := handle.NewManage()
	manage.AddBeforeHandle(0, &TestHandle{})
	manage.AddBeforeHandle(1, &TestHandle{})
	clusterHandle := newClusterHandle("127.0.0.1:8000")
	go func() {
		clusterHandle.Cluster.ConnectToServer()
	}()
	manage.AddClusterHandle(clusterHandle)
	ws.SetManage(manage)
	ws.OnInit(9000)
	ws.StartServer()
}
