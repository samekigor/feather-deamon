package socket

import (
	"log"
	"net"
	"os"

	"github.com/samekigor/feather-deamon/internal/config"
	"github.com/samekigor/feather-deamon/shared/proto"
	"google.golang.org/grpc"
)

type SocketDetails struct {
	Path string
}

type Server struct {
	// proto.UnimplementedFeatherServer
	proto.UnimplementedFeatherServer
}

var socketDetails SocketDetails

func initSocketDetails() {
	socketDetails.Path = config.GetValueFromConfig("socket.path")
}

func Cleanup() {
	log.Println("Cleaning up resources...")
	if err := os.Remove(socketDetails.Path); err != nil {
		log.Printf("Failed to remove socket file: %v", err)
	}
}

func StartGRPCServer() {
	initSocketDetails()
	lis, err := net.Listen("unix", socketDetails.Path)
	if err != nil {
		log.Fatalf("Failed to start Unix socket listener: %v", err)
	}
	defer lis.Close()

	if err := os.Chmod(socketDetails.Path, 0755); err != nil {
		log.Fatalf("Failed to set permissions on socket file: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterFeatherServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("Unix socket server started at %s", socketDetails.Path)
}
