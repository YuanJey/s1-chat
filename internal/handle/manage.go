package handle

import (
	"s1-chat/pkg/structs"
)

type Manage struct {
	beforeHandles    []Handle
	inProcessHandles []Handle
	afterHandles     []Handle
}

func NewManage() *Manage {
	var beforeHandles []Handle
	var inProcessHandles []Handle
	var afterHandles []Handle
	return &Manage{
		afterHandles:     afterHandles,
		inProcessHandles: inProcessHandles,
		beforeHandles:    beforeHandles,
	}
}
func (m *Manage) AddBeforeHandle(handle Handle) {
	m.beforeHandles = append(m.beforeHandles, handle)
}
func (m *Manage) AddInProcessHandle(handle Handle) {
	m.beforeHandles = append(m.beforeHandles, handle)
}
func (m *Manage) AddAfterHandle(handle Handle) {
	m.beforeHandles = append(m.beforeHandles, handle)
}
func (m *Manage) Work(msg *structs.Message) {
	for i := range m.beforeHandles {
		m.beforeHandles[i].Processing(msg)
	}
	for i := range m.inProcessHandles {
		m.inProcessHandles[i].Processing(msg)
	}
	for i := range m.afterHandles {
		m.afterHandles[i].Processing(msg)
	}
}
