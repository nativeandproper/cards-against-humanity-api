package main

import (
	"fmt"
	"os"
)

// GetEnvOrPanic panics if no env variable is set
func getEnvOrPanic(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	panic(fmt.Sprintf("env variable not found: %s", key))
}
