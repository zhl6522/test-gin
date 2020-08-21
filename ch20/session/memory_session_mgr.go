package session

import (
	"fmt"
	"sync"
	uuid "github.com/satori/go.uuid"
)

// 定义对象
type MemorySessionMgr struct {
	sessionMap	map[string]Session
	rwlock		sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() SessionMgr {
	sr := &MemorySessionMgr{
		sessionMap:make(map[string]Session, 1024),
	}
	return sr
}

func (m *MemorySessionMgr)Init(addr string,options ...string)(err error) {
	return
}

func (m *MemorySessionMgr)CreateSession()(session Session, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// go get github.com/satori/go.uuid
	// 用uuid作为sessionId
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	// 转string
	sessionId := id.String()
	// 创建一个session
	session = NewMemorySession(sessionId)

	return
}
