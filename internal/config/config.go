package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppEnv string
	Port   int
}

// Load loads configuration from environment variables.
// - APP_ENV (default: local)
// - PORT    (default: 8080)
func Load() Config {
	return Config{
		AppEnv: getEnv("APP_ENV", "local"),
		Port:   getEnvInt("PORT", 8080),
	}
}

func (c Config) Addr() string {
	return ":" + strconv.Itoa(c.Port)
}

func getEnv(key, def string) string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return def
	}
	return v
}

func getEnvInt(key string, def int) int {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
