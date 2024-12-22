package main

import (
	"fmt"
	"time"

	"github.com/samekigor/feather-deamon/internal"
)

func main() {
	internal.LoadConfigFile()
	internal.SetupLogging()

	for {
		fmt.Println("Running...")
		time.Sleep(1 * time.Second)
	}

}
