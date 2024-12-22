package socket

import (
	"log"
	"net"
	"strings"
)

func parseMessage(msg string, conn net.Conn) {
	if strings.HasPrefix(msg, "start") {
		command := strings.Split(msg, " ")
		if len(strings.Split(msg, " ")) != 2 {
			log.Printf("Invalid message: %s", msg)
			return
		}
		_, err := conn.Write([]byte("starting container " + command[1] + "\n"))
		if err != nil {
			log.Printf("Failed to write response: %v", err)
			return
		}
	}
	if strings.HasPrefix(msg, "stop") {
		command := strings.Split(msg, " ")
		if len(strings.Split(msg, " ")) != 2 {
			log.Printf("Invalid message: %s", msg)
			return
		}
		_, err := conn.Write([]byte("stopping container " + command[1] + "\n"))
		if err != nil {
			log.Printf("Failed to write response: %v", err)
			return
		}
	}
}
