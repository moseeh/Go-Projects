package client

import (
	"errors"
	"net"

	"net-cat/models"
)

// AddClient attempts to add a new client to the Clients map.
// It returns an error if the name is already in use.
func AddClient(name string, conn net.Conn) error {
	models.ClientsMutex.Lock()
	defer models.ClientsMutex.Unlock()

	if _, exists := models.Clients[name]; exists {
		return errors.New("name already in use")
	}

	models.Clients[name] = conn
	BroadcastJoin(name)
	return nil
}

// RemoveClient removes a client from the Clients map.
func RemoveClient(name string) {
	models.ClientsMutex.Lock()
	defer models.ClientsMutex.Unlock()

	if _, exists := models.Clients[name]; exists {
		delete(models.Clients, name)
		BroadcastLeave(name)
	}
}
