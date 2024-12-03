package models

import (
	"net"
	"sync"
)

var (
	MaxConnections = 10
	Clients        = make(map[string]net.Conn)
	ClientsMutex   sync.Mutex
)

type Server struct {
	Prevchat []string
	Mutex    sync.Mutex
	Logo     []byte
}
