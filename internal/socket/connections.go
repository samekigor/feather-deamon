package socket

import (
	"bufio"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("Connection established: %v", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		log.Printf("Received message: %s", msg)
		parseMessage(msg, conn)
		_, err := conn.Write([]byte("Received: " + msg + "\n"))
		if err != nil {
			log.Printf("Failed to write response: %v", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection: %v", err)
	}

	log.Printf("Connection closed: %v", conn.RemoteAddr())
}

func HandleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	log.Println("Daemon shutting down")
	Cleanup()
	os.Exit(0)
}
