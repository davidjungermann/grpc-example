package utils

import (
	"errors"
	"os"
)

func GetEnvironmentVariable(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("environment variable " + key + " is not set")
	}
	return value, nil
}

func GetPortFromEnv(key string, defaultVal string) string {
	val, err := GetEnvironmentVariable(key)
	if err != nil {
		val = defaultVal
	}
	return ":" + val
}
