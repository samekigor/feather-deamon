package logger

import (
	"log"
	"os"

	"github.com/samekigor/feather-deamon/internal/config"
)

type LoggerDetails struct {
	filePath string
	logFile  os.File
}

var loggerDetails LoggerDetails

func SetupLogging() {
	filePath := config.GetValueFromConfig("logger.path")
	if filePath == "" {
		log.Panicf("filepaths.logs not defined in config yaml")
	}
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)

	loggerDetails.filePath = filePath
	loggerDetails.logFile = *logFile

}

func CloseLogging() {
	log.Printf("Closeing logging file: %s", loggerDetails.filePath)
	loggerDetails.logFile.Close()
}
