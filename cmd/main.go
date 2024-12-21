package main

import (
	"github.com/samekigor/feather-deamon/internal"
)

func main() {
	internal.InitEnv()
	internal.SetupLogging()
}
