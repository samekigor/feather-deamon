package internal

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitEnv() {
	viper.SetEnvPrefix("FEATHER")
	viper.AutomaticEnv()
}

func LoadConfigFile(configPath string) {
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
