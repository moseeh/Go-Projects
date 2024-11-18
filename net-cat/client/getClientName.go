package client

import (
	"bufio"
	"net"
	"strings"
)

func GetClientName(conn net.Conn) (string, error) {
	conn.Write([]byte("[ENTER YOUR NAME]: "))
	// Read the client's name from the connection
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(name), nil
}
