package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"net-cat/client"
	"net-cat/models"
	"net-cat/utils"
)

type newServer struct {
	models.Server
}

func Server(addr string) {
	server := newServer{
		Server: models.Server{
			Prevchat: make([]string, 0),
		},
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("error listening", err)
	}
	fmt.Printf("Listening on port %s  ...\n", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connections")
			continue
		}
		fmt.Println(len(models.Clients))
		go server.handleConnection(conn)
	}
}

func (s *newServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection from", conn.RemoteAddr())
	conn.Write(utils.ReadLogo())
	name, err := client.GetClientName(conn)
	if err != nil {
		fmt.Println("Error getting client name:", err)
		return
	}

	// Add the client to the chat
	client.AddClient(name, conn)
	defer client.RemoveClient(name)

	// Handle messages from the client
	s.handleMessages(name, conn)
}

func (s *newServer) handleMessages(name string, conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading message:", err)
			}
			return
		}

		if strings.TrimSpace(msg) == "" {
			continue
		}

		broadcastMessage(name, msg)
	}
}

func broadcastMessage(name, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMsg := fmt.Sprintf("[%s][%s]: %s", timestamp, name, msg)

	models.ClientsMutex.Lock()
	defer models.ClientsMutex.Unlock()
	for _, c := range models.Clients {
		fmt.Fprint(c, formattedMsg)
	}
}
