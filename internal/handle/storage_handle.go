package handle

import (
	"s1-chat/pkg/structs"
)

type StorageHandle struct {
	Name     string
	msgTypes []int
}

// Processing 存储消息
func (s *StorageHandle) Processing(msg *structs.Message) bool {
	//TODO implement me
	panic("implement me")
}
