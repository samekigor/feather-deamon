package internal

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetEnv(envVar string) (envVal string) {
	viper.SetEnvPrefix("FEATHER")
	viper.AutomaticEnv()
	envVal = viper.GetString(envVar)
	if envVal == "" {
		log.Fatalf("Env variable={%s} not found", envVar)
	}
	return envVal
}

func LoadConfigFile() {
	configPath := GetEnv("CONFIG_FILE_PATH")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file: %v", err))
	}
}

func GetValueFromConfig(yamlKey string) string {
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file: %v", err))
	}

	logsPath := viper.GetString(yamlKey)
	if logsPath == "" {
		panic(fmt.Sprintf("Key={%s} not found", yamlKey))
	}

	return logsPath
}
