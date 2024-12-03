package client_test

import (
	"bytes"
	"net"
	"strings"
	"testing"

	"net-cat/client"
	"net-cat/models"
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

// Test AddClient
func TestAddClient(t *testing.T) {
	mockConn := NewMockConn()
	models.Clients = make(map[string]net.Conn) // Reset clients

	client.AddClient("TestClient", mockConn)

	if _, exists := models.Clients["TestClient"]; !exists {
		t.Errorf("Expected client 'TestClient' to be added")
	}

	output := mockConn.WriteBuffer.String()
	if strings.Contains(output, "client already exists") {
		t.Errorf("Unexpected error message: %s", output)
	}
}

// Test AddClient with existing name
func TestAddClientWithExistingName(t *testing.T) {
	mockConn1 := NewMockConn()
	mockConn2 := NewMockConn()

	models.Clients = make(map[string]net.Conn) // Reset clients
	client.AddClient("TestClient", mockConn1)

	client.AddClient("TestClient", mockConn2)

	output := mockConn2.WriteBuffer.String()
	if !strings.Contains(output, "client already exists") {
		t.Errorf("Expected error message for duplicate name, got: %s", output)
	}
}

// Test RemoveClient
func TestRemoveClient(t *testing.T) {
	mockConn := NewMockConn()

	models.Clients = make(map[string]net.Conn) // Reset clients
	client.AddClient("TestClient", mockConn)

	client.RemoveClient("TestClient")

	if _, exists := models.Clients["TestClient"]; exists {
		t.Errorf("Expected client 'TestClient' to be removed")
	}
}

// Test BroadcastJoin
func TestBroadcastJoin(t *testing.T) {
	mockConn1 := NewMockConn()
	mockConn2 := NewMockConn()

	models.Clients = make(map[string]net.Conn) // Reset clients
	client.AddClient("User1", mockConn1)
	client.AddClient("User2", mockConn2)

	client.BroadcastJoin("User2")

	output := mockConn1.WriteBuffer.String()
	if !strings.Contains(output, "User2 has joined our chat") {
		t.Errorf("Expected join message, got: %s", output)
	}
}

// Test BroadcastLeave
func TestBroadcastLeave(t *testing.T) {
	mockConn1 := NewMockConn()
	mockConn2 := NewMockConn()

	models.Clients = make(map[string]net.Conn) // Reset clients
	client.AddClient("User1", mockConn1)
	client.AddClient("User2", mockConn2)

	client.BroadcastLeave("User2")

	output := mockConn1.WriteBuffer.String()
	if !strings.Contains(output, "User2 has left our chat") {
		t.Errorf("Expected leave message, got: %s", output)
	}
}

// Test GetClientName
func TestGetClientName(t *testing.T) {
	mockConn := NewMockConn()
	mockConn.ReadBuffer.WriteString("UniqueName\n") // Simulate name input

	models.Clients = make(map[string]net.Conn) // Reset clients

	name, err := client.GetClientName(mockConn)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if name != "UniqueName" {
		t.Errorf("Expected name 'UniqueName', got: %s", name)
	}
}

