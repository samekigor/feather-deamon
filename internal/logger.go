package internal

import (
	"log"
	"os"
)

func SetupLogging() {
	filePath := GetValueFromConfig("filepaths.logs")
	if filePath == "" {
		log.Panicf("filepaths.logs not defined in config yaml")
	}
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)

	defer logFile.Close()
}
