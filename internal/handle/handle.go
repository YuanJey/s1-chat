package handle

import (
	"s1-chat/pkg/structs"
)

type Handle interface {
	Processing(msg *structs.Message)
}
