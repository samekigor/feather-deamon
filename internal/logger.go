package internal

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func SetupLogging() {
	filePath := viper.GetString("LOG_FILE_PATH")
	if filePath == "" {
		log.Panicf("LOG_FILE_PATH is not set in configuration")
	}
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)

	defer logFile.Close()
}
