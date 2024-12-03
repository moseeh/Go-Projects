package client

import (
	"net"

	"net-cat/models"
)

func AddClient(name string, conn net.Conn) {
	models.ClientsMutex.Lock()
	defer models.ClientsMutex.Unlock()

	if _, exists := models.Clients[name]; exists {
		conn.Write([]byte("client already exists...reconnect with different name"))
		conn.Close()
		return
	}

	models.Clients[name] = conn
	BroadcastJoin(name)
}

func RemoveClient(name string) {
	models.ClientsMutex.Lock()
	defer models.ClientsMutex.Unlock()
	delete(models.Clients, name)
	BroadcastLeave(name)
}
