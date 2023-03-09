package kop

import (
	"net"
	"sync"
)

// NetworkContext
// authed 记录Kafka鉴权状态
type NetworkContext struct {
	ctxMutex sync.RWMutex
	authed   bool
	Addr     net.Addr
}

func (n *NetworkContext) Authed(authed bool) {
	n.ctxMutex.RLock()
	n.authed = authed
	n.ctxMutex.RUnlock()
}

func (n *NetworkContext) IsAuthed() bool {
	n.ctxMutex.RLock()
	defer n.ctxMutex.RUnlock()
	return n.authed
}

func getContext(conn net.Conn) *NetworkContext {
	return &NetworkContext{
		Addr: conn.RemoteAddr(),
	}
}
