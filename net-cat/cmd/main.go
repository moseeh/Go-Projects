package main

import (
	"fmt"
	"os"

	"net-cat/server"
)

func main() {
	args := os.Args[1:]
	var port string
	if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(args) == 1 {
		port = ":" + args[0]
	} else {
		port = ":8989"
	}
	server.Server(port)
}
