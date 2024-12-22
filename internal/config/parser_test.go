package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestGetValueFromConfig(t *testing.T) {
	configContent := `filepaths:
  logs: "/var/log/test.log"
  socket: "/var/run/test.sock"`
	tempFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(configContent)
	if err != nil {
		t.Fatalf("Failed to write to temp config file: %v", err)
	}
	tempFile.Close()

	viper.SetConfigFile(tempFile.Name())

	expectedValue := "/var/log/test.log"
	actualValue := GetValueFromConfig("filepaths.logs")
	if actualValue != expectedValue {
		t.Errorf("Expected %s, but got %s", expectedValue, actualValue)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic for missing key, but function did not panic")
		} else if r != "Key={filepaths.invalid} not found" {
			t.Errorf("Unexpected panic message: %v", r)
		}
	}()

	GetValueFromConfig("filepaths.invalid")
}
