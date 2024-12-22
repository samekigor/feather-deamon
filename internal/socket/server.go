package socket

import (
	"log"
	"net"
	"os"

	"github.com/samekigor/feather-deamon/internal/config"
)

type SocketDetails struct {
	Path string
}

var socketDetails SocketDetails

func initSocketDetails() {
	socketDetails.Path = config.GetValueFromConfig("socket.path")
}

func StartUnixSocketServer() {
	initSocketDetails()
	listener, err := net.Listen("unix", socketDetails.Path)
	if err != nil {
		log.Fatalf("Failed to start Unix socket listener: %v", err)
	}
	defer listener.Close()

	if err := os.Chmod(socketDetails.Path, 0755); err != nil {
		log.Fatalf("Failed to set permissions on socket file: %v", err)
	}

	log.Printf("Unix socket server started at %s", socketDetails.Path)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go HandleConnection(conn)
	}
}

func Cleanup() {
	log.Println("Cleaning up resources...")
	if err := os.Remove(socketDetails.Path); err != nil {
		log.Printf("Failed to remove socket file: %v", err)
	}
}
