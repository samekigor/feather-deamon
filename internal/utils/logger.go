package logger

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func SetupLogging() {
	viper.SetEnvPrefix("FEATHER")
	viper.AutomaticEnv()

	filePath := viper.GetString("LOG_FILE_PATH")
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(logFile)
}
