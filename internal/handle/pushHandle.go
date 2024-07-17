package handle

import (
	"s1-chat/pkg/structs"
)

type PushHandle struct {
	Name     string
	msgTypes []int
}

func (p *PushHandle) Processing(msg *structs.Message) {
	//TODO implement me
	panic("implement me")
}
