package server_test

import (
	"bytes"
	"net"
	"strings"
	"sync"
	"testing"
	"time"

	"net-cat/client"
	"net-cat/models"
	"net-cat/server"
	"net-cat/utils"
)

// Mock connection for testing
type MockConn struct {
	net.Conn
	ReadBuffer  *bytes.Buffer
	WriteBuffer *bytes.Buffer
}

func NewMockConn() *MockConn {
	return &MockConn{
		ReadBuffer:  new(bytes.Buffer),
		WriteBuffer: new(bytes.Buffer),
	}
}

func (c *MockConn) Read(p []byte) (n int, err error) {
	return c.ReadBuffer.Read(p)
}

func (c *MockConn) Write(p []byte) (n int, err error) {
	return c.WriteBuffer.Write(p)
}

func (c *MockConn) Close() error {
	return nil
}

// Test Server Initialization
func TestServerInitialization(t *testing.T) {
	addr := ":8080"
	go func() {
		server.Server(addr)
	}()

	time.Sleep(100 * time.Millisecond) // Allow some time for server to start
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	conn.Close()
}

// Test handling a new connection
func TestHandleConnection(t *testing.T) {
	mockConn := NewMockConn()
	mockConn.ReadBuffer.WriteString("TestClient\n") // Simulate client name input

	testServer := &server.NewServer{
		Server: models.Server{
			Prevchat: make([]string, 0),
			Mutex:    sync.Mutex{},
		},
	}

	testServer.Logo = utils.ReadLogo()

	go testServer.HandleConnection(mockConn)

	time.Sleep(50 * time.Millisecond)

	output := mockConn.WriteBuffer.String()
	if !strings.Contains(output, string(testServer.Logo)) {
		t.Errorf("Expected welcome message, got %s", output)
	}
}

// Test broadcasting messages
func TestBroadcastMessage(t *testing.T) {
	mockConn1 := NewMockConn()
	mockConn2 := NewMockConn()

	models.Clients = make(map[string]net.Conn)
	client.AddClient("User1", mockConn1)
	client.AddClient("User2", mockConn2)

	testServer := &server.NewServer{
		Server: models.Server{
			Prevchat: make([]string, 0),
			Mutex:    sync.Mutex{},
		},
	}

	testServer.BroadcastMessage("User1", "Hello everyone!\n")

	output1 := mockConn1.WriteBuffer.String()
	output2 := mockConn2.WriteBuffer.String()

	if !strings.Contains(output1, "[User1]:Hello everyone!") {
		t.Errorf("Expected broadcast message for User1, got %s", output1)
	}
	if !strings.Contains(output2, "[User1]:Hello everyone!") {
		t.Errorf("Expected broadcast message for User2, got %s", output2)
	}
}

// Test handling messages
func TestHandleMessages(t *testing.T) {
	mockConn := NewMockConn()
	mockConn.ReadBuffer.WriteString("Hello there!\n")

	models.Clients = make(map[string]net.Conn)
	client.AddClient("User1", mockConn)

	testServer := &server.NewServer{
		Server: models.Server{
			Prevchat: make([]string, 0),
			Mutex:    sync.Mutex{},
		},
	}

	go testServer.HandleMessages("User1", mockConn)

	time.Sleep(50 * time.Millisecond)

	output := mockConn.WriteBuffer.String()
	if !strings.Contains(output, "[User1]:Hello there!") {
		t.Errorf("Expected handled message, got %s", output)
	}
}
