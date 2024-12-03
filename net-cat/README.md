# Net-Cat: TCP-Based Group Chat Application

A TCP-based chat server implementation in Go that enables multiple clients to communicate in a group chat environment. This project recreates core functionalities of the NetCat (`nc`) command-line utility, focusing on TCP connections and real-time message broadcasting.

## Features

- TCP server supporting multiple concurrent client connections
- Real-time message broadcasting to all connected clients
- User authentication with name requirement
- Connection limit management (maximum 10 concurrent users)
- Automatic message timestamping and user identification
- Message history preservation and replay for new clients
- Join/Leave notifications for all connected users
- Graceful connection handling and error management

## Technical Requirements

- Go (1.15 or later recommended)
- TCP/IP network connection
- Unix-like operating system (Linux, macOS) or Windows

## Installation

1. Clone the repository:
```bash
git clone https://learn.zone01kisumu.ke/git/sjuma/net-cat.git
cd net-cat
```

2. Build the application using either method:
```bash
# Method 1: Using build script
./build.sh

# Method 2: 
go build -o TCPChat ./cmd
```
This will create a binary named `TCPChat` in your current directory.

## Usage

### Starting the Server

```bash
# Using default port (8989)
go run ./cmd
#OR
./TCPChat

# Using custom port
go run ./cmd <port-number>
#OR
./TCPChat <port-number>
```

### Connecting as a Client

You can connect to the server using the standard `nc` command or any TCP client:

```bash
nc localhost <port-number>
```

Upon connection, you will be prompted to enter your name before joining the chat.

## Message Format

Messages in the chat are formatted as follows:
```
[YYYY-MM-DD HH:MM:SS][username]: message
```

Example:
```
[2024-01-20 15:48:41][John]: Hello, everyone!
```

## Technical Implementation Details

### Core Components

- **TCP Server**: Handles incoming connections and maintains client sessions
- **Client Handler**: Manages individual client connections and message broadcasting
- **Message Broadcasting**: Implements concurrent message delivery to all connected clients
- **Connection Management**: Controls the number of concurrent connections (max 10)
- **Error Handling**: Robust error management for both server and client operations

### Concurrency Features

- Goroutines for handling multiple client connections
- Channels for message broadcasting and synchronization
- Mutex implementation for thread-safe operations
- Connection pooling and management

## Error Handling

The application handles various error scenarios:
- Connection failures
- Network interruptions
- Invalid client input
- Server capacity limitations
- Unexpected client disconnections

## Project Structure

```
net-cat/
├── cmd/
│   └── main.go        # Application entry point
├── server/            # Server implementation
├── client/            # Client handling logic
├── utils/             # Shared utilities         
├── build.sh           # Build script
├── models             # structs and variables
└── README.md          # This file
```

## Testing

The project includes test files for both server and client components. Run tests using:

```bash
go test ./...
```

## Limitations

- Maximum 10 concurrent connections
- Server must be running before clients can connect
- Messages are not persisted between server restarts

## Contributing

Feel free to open issues or submit pull requests for any improvements or bug fixes.

## License

[Add your license information here]
