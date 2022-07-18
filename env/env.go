package env

import "os"

// Getenv return the defined environment variable and provide a fallback if the environment variable is not defined
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
