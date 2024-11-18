package client

import (
	"net"

	"net-cat/models"
)

func AddClient(name string, conn net.Conn) {
	models.ClientsMutex.Lock()
	defer models.ClientsMutex.Unlock()
	models.Clients[name] = conn
	BroadcastJoin(name)
}

func RemoveClient(name string) {
	models.ClientsMutex.Lock()
	defer models.ClientsMutex.Unlock()
	delete(models.Clients, name)
	BroadcastLeave(name)
}
