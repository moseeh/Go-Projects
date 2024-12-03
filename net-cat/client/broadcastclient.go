package client

import (
	"fmt"

	"net-cat/models"
)

func BroadcastJoin(name string) {
	// Notify all other models.Clients that a new client has joined
	for n, c := range models.Clients {
		if n != name {
			fmt.Fprintf(c, "%s has joined our chat\n", name)
		}
	}
}

func BroadcastLeave(name string) {
	// Notify all other models.Clients that a client has left
	for n, c := range models.Clients {
		if n != name {
			fmt.Fprintf(c, "%s has left our chat\n", name)
		}
	}
}
