package handle

import (
	"s1-chat/pkg/structs"
)

type Manage struct {
	beforeHandlesMaxCount    int
	inProcessHandlesMaxCount int
	afterHandlesCountMax     int
	beforeHandles            map[int]Handle
	inProcessHandles         map[int]Handle
	afterHandles             map[int]Handle
	clusterHandles           []Handle
}

func NewManage() *Manage {
	return &Manage{
		afterHandles:     make(map[int]Handle),
		inProcessHandles: make(map[int]Handle),
		beforeHandles:    make(map[int]Handle),
		clusterHandles:   []Handle{},
	}
}
func (m *Manage) AddBeforeHandle(index int, handle Handle) {
	m.beforeHandles[index] = handle
	if m.beforeHandlesMaxCount < index {
		m.beforeHandlesMaxCount = index
		return
	}
	m.beforeHandlesMaxCount = index
}
func (m *Manage) AddInProcessHandle(index int, handle Handle) {
	m.inProcessHandles[index] = handle
	if m.inProcessHandlesMaxCount < index {
		m.inProcessHandlesMaxCount = index
		return
	}
	m.inProcessHandlesMaxCount = index
}
func (m *Manage) AddAfterHandle(index int, handle Handle) {
	m.afterHandles[index] = handle
	if m.afterHandlesCountMax < index {
		m.afterHandlesCountMax = index
		return
	}
	m.afterHandlesCountMax = index
}
func (m *Manage) AddClusterHandle(handle Handle) {
	m.clusterHandles = append(m.clusterHandles, handle)
}

func (m *Manage) ProcessMessage(msg *structs.Message) {
	for i := range m.beforeHandlesMaxCount + 1 {
		if handle, ok := m.beforeHandles[i]; ok {
			handle.Processing(msg)
		}
	}
	for i := range m.inProcessHandlesMaxCount + 1 {
		if handle, ok := m.inProcessHandles[i]; ok {
			handle.Processing(msg)
		}
	}
	for i := range m.afterHandlesCountMax + 1 {
		if handle, ok := m.afterHandles[i]; ok {
			handle.Processing(msg)
		}
	}
	for i := range m.clusterHandles {
		m.clusterHandles[i].Processing(msg)
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
