package handle

import (
	"s1-chat/pkg/structs"
)

type Manage struct {
	beforeHandlesCount    int
	inProcessHandlesCount int
	afterHandlesCount     int
	beforeHandles         map[int]Handle
	inProcessHandles      map[int]Handle
	afterHandles          map[int]Handle
}

func NewManage() *Manage {
	return &Manage{
		afterHandles:     make(map[int]Handle),
		inProcessHandles: make(map[int]Handle),
		beforeHandles:    make(map[int]Handle),
	}
}
func (m *Manage) AddBeforeHandle(index int, handle Handle) {
	m.beforeHandles[index] = handle
}
func (m *Manage) AddInProcessHandle(index int, handle Handle) {
	m.inProcessHandles[index] = handle
}
func (m *Manage) AddAfterHandle(index int, handle Handle) {
	m.afterHandles[index] = handle
}

func (m *Manage) ProcessMessage(msg *structs.Message) {
	for i := range m.beforeHandlesCount {
		if handle, ok := m.beforeHandles[i]; ok {
			handle.Processing(msg)
		}
	}
	for i := range m.inProcessHandlesCount {
		if handle, ok := m.inProcessHandles[i]; ok {
			handle.Processing(msg)
		}
	}
	for i := range m.afterHandlesCount {
		if handle, ok := m.afterHandles[i]; ok {
			handle.Processing(msg)
		}
	}
}

type Manage2 struct {
	handles []Handle
}

func NewManage2() *Manage2 {
	var handles []Handle
	return &Manage2{handles: handles}
}
func (m *Manage2) AddHandle(handle Handle) {
	m.handles = append(m.handles, handle)
}

// ProcessMessage bool 是否处理完成
func (m *Manage2) ProcessMessage(msg *structs.Message) bool {
	for i := range m.handles {
		isContinue := m.handles[i].Processing(msg)
		if isContinue {
			continue
		} else {
			return true
		}
	}
	return true
}
