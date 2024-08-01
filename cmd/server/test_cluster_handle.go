package main

import (
	"s1-chat/internal/client"
	"s1-chat/pkg/structs"
	"s1-chat/pkg/utils"
)

type ClusterHandle struct {
	Cluster *client.TCPClient
}

func (c *ClusterHandle) Processing(msg *structs.Message) bool {
	c.Cluster.SendMsg(utils.StructToJsonString(msg))
	return true
}

func newClusterHandle(addr string) *ClusterHandle {
	tcpClient := client.NewTCPClient(addr)
	return &ClusterHandle{Cluster: tcpClient}
}
func (c *ClusterHandle) ReConn() {
	c.Cluster.ConnectToServer()
}
