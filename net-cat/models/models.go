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

type Client struct {
	Conn     net.Conn
	Username string
	Messages chan string
}

type Server struct {
	Messages Message
	Prevchat []string
	Mutex    sync.Mutex
}

type Message struct {
	Text       string
	Senderconn net.Conn
}
