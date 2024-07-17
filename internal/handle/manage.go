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

func (m *Manage) ProcessMessage(msg *structs.Message) {
	for i := range m.beforeHandles {
		isContinue := m.beforeHandles[i].Processing(msg)
		if isContinue {
			continue
		} else {
			return
		}
	}
	for i := range m.inProcessHandles {
		isContinue := m.inProcessHandles[i].Processing(msg)
		if isContinue {
			continue
		} else {
			return
		}
	}
	for i := range m.afterHandles {
		isContinue := m.afterHandles[i].Processing(msg)
		if isContinue {
			continue
		} else {
			return
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
