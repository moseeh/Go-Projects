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

type NewServer struct {
	models.Server
}

func Server(addr string) {
	server := NewServer{
		Server: models.Server{
			Prevchat: make([]string, 0),
		},
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	fmt.Printf("Server listening on port %s...\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New connection from", conn.RemoteAddr())

		go server.HandleConnection(conn)
	}
}

func (s *NewServer) HandleConnection(conn net.Conn) {
	defer conn.Close()
	s.Mutex.Lock()

	if len(models.Clients) >= models.MaxConnections {
		conn.Write([]byte("Error: Maximum number of connections reached. Please try again later.\n"))
		conn.Close()
		s.Mutex.Unlock()
		return
	}

	s.Mutex.Unlock()
	s.Logo = utils.ReadLogo()
	conn.Write(s.Logo)

	name, err := client.GetClientName(conn)
	if err != nil {
		fmt.Println("Error getting client name:", err)
		conn.Write([]byte("Error: Could not get your name. Please try again.\n"))
		conn.Close()
		return
	}

	client.AddClient(name, conn)

	defer client.RemoveClient(name)

	s.SendPreviousChats(conn)

	s.HandleMessages(name, conn)
}

func (s *NewServer) SendPreviousChats(conn net.Conn) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if len(s.Prevchat) > 0 {
		for _, msg := range s.Prevchat {
			conn.Write([]byte(msg))
		}
	}
}

func (s *NewServer) HandleMessages(name string, conn net.Conn) {
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
			timestamp := time.Now().Format("2006-01-02 15:04:05")

			formattedMsg := fmt.Sprintf("[%s][%s]:%s", timestamp, name, msg)
			fmt.Fprint(conn, "\033[F\033[K")
			fmt.Fprint(conn, formattedMsg)
			continue
		}

		s.BroadcastMessage(name, msg)
	}
}

func (s *NewServer) BroadcastMessage(name, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	formattedMsg := fmt.Sprintf("[%s][%s]:%s", timestamp, name, msg)

	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Prevchat = append(s.Prevchat, formattedMsg)

	for n, c := range models.Clients {
		if n == name {
			fmt.Fprint(c, "\033[F\033[K")
		}
		fmt.Fprint(c, formattedMsg)
	}
}
