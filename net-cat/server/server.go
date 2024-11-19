package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"

	"net-cat/client"
	"net-cat/models"
	"net-cat/utils"
)

type newServer struct {
	models.Server
	mu sync.Mutex // Mutex to ensure thread-safe access to Prevchat
}

func Server(addr string) {
	server := newServer{
		Server: models.Server{
			Prevchat: make([]string, 0),
		},
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Listening on port %s...\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go server.handleConnection(conn)
	}
}

func (s *newServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Check if the chat is full
	models.ClientsMutex.Lock()
	if len(models.Clients) >= models.MaxConnections {
		models.ClientsMutex.Unlock()
		conn.Write([]byte("Chat room is full, try again later.\n"))
		return
	}
	models.ClientsMutex.Unlock()

	fmt.Println("New connection from", conn.RemoteAddr())

	// Send welcome logo
	conn.Write(utils.ReadLogo())

	// Get client name
	name, err := client.GetClientName(conn)
	if err != nil {
		fmt.Println("Error getting client name:", err)
		conn.Write([]byte("Failed to retrieve your name. Disconnecting.\n"))
		return
	}

	// Send previous chat history
	s.mu.Lock()
	for _, v := range s.Prevchat {
		conn.Write([]byte(v + "\n"))
	}
	s.mu.Unlock()

	// Add client to the chat
	client.AddClient(name, conn)
	defer client.RemoveClient(name)

	fmt.Printf("%s has joined the chat.\n", name)

	// Handle messages from the client
	s.handleMessages(name, conn)
}

func (s *newServer) handleMessages(name string, conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading message from %s: %v\n", name, err)
			}
			fmt.Printf("%s has left the chat.\n", name)
			return
		}

		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}

		// Format the message with timestamp
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formattedMsg := fmt.Sprintf("[%s][%s]: %s", timestamp, name, msg)

		// Append the message to the chat history
		s.mu.Lock()
		s.Prevchat = append(s.Prevchat, formattedMsg)
		s.mu.Unlock()

		// Broadcast the message
		broadcastMessage(formattedMsg)
	}
}

func broadcastMessage(msg string) {
	models.ClientsMutex.Lock()
	defer models.ClientsMutex.Unlock()

	for name, c := range models.Clients {
		_, err := fmt.Fprintln(c, msg)
		if err != nil {
			fmt.Printf("Error sending message to %s: %v\n", name, err)
		}
	}
}
