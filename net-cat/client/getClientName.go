package client

import (
	"bufio"
	"net"
	"net-cat/models"
	"strings"
)

func GetClientName(conn net.Conn) (string, error) {

	for {
		conn.Write([]byte("[ENTER YOUR NAME]: "))
		reader := bufio.NewReader(conn)

		name, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		name = strings.TrimSpace(name)
		models.ClientsMutex.Lock()
		_, exists := models.Clients[name]
		models.ClientsMutex.Unlock()

		if !exists {
			return name, nil
		}
		conn.Write([]byte("name already taken....try again.\n"))

	}
}
