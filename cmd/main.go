package main

import (
	"log"
	"time"

	"github.com/samekigor/feather-deamon/internal/config"
	"github.com/samekigor/feather-deamon/internal/logger"
	"github.com/samekigor/feather-deamon/internal/socket"
)

func main() {
	config.LoadConfigFile()
	logger.SetupLogging()
	defer logger.CloseLogging()

	go socket.StartUnixSocketServer()

	go socket.HandleSignals()
	for {
		log.Println("Running...")
		time.Sleep(1 * time.Second)
	}

}
