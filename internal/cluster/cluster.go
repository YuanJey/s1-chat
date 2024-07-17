package cluster

import "s1-chat/pkg/structs"

type Cluster struct {
	registerServer *RegisterServer
}

func (r *Cluster) SendMsg(msg structs.Message) {

}
