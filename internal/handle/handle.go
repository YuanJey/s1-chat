package handle

import (
	"s1-chat/pkg/structs"
)

// Handle false 结束 true 继续
type Handle interface {
	Processing(msg *structs.Message) bool
}
