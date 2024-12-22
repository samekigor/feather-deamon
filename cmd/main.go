package main

import (
	"fmt"
	"time"

	"github.com/samekigor/feather-deamon/internal/config"
	"github.com/samekigor/feather-deamon/internal/logger"
)

func main() {
	config.LoadConfigFile()
	logger.SetupLogging()
	defer logger.CloseLogging()
	for {
		fmt.Println("Running...")
		time.Sleep(1 * time.Second)
	}

}
