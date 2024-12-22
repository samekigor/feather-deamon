package logger

import (
	"log"
	"os"

	"github.com/samekigor/feather-deamon/internal/config"
)

type Logger struct {
	filePath string
	logFile  os.File
}

var loggerInfo Logger

func SetupLogging() {
	filePath := config.GetValueFromConfig("filepaths.logs")
	if filePath == "" {
		log.Panicf("filepaths.logs not defined in config yaml")
	}
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0622)
	if err != nil {
		log.Panicf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)

	loggerInfo.filePath = filePath
	loggerInfo.logFile = *logFile

}

func CloseLogging() {
	log.Printf("Closeing logging file: %s", loggerInfo.filePath)
	loggerInfo.logFile.Close()
}
