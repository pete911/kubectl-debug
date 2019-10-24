package util

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func GetHomeDir() (string, error) {

	u, err := user.Current()
	if err != nil {
		if userHome := os.Getenv("HOME"); userHome != "" {
			return userHome, nil
		}
		return "", fmt.Errorf("cannot get user's home directory: %v", err)
	}
	return u.HomeDir, nil
}

// GetBoolEnvVar gets string environment variable by envVarName. Returns defaultValue if empty.
func GetStringEnvVar(envVarName, defaultValue string) string {
	if value := os.Getenv(envVarName); value != "" {
		return value
	}
	return defaultValue
}

// GetBoolEnvVar gets bool environment variable by envVarName. Returns defaultValue if empty.
func GetBoolEnvVar(envVarName string, defaultValue bool) bool {
	value := os.Getenv(envVarName)
	if value == "" {
		return defaultValue
	}

	if boolValue, err := strconv.ParseBool(value); err == nil {
		return boolValue
	}
	return defaultValue
}
