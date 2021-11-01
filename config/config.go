package config

import "os"

// Config func to get env value
func GetEnv(key string) string {
	return os.Getenv(key)
}
