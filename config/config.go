package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostServiceNewHost string
	PostServiceNewPort int
	HttpPort           string
	LogLevel           string
}

func Load() Config {
	c := Config{}

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":9002"))
	c.PostServiceNewHost = cast.ToString(getOrReturnDefault("POST_SERVICE_NEW_HOST", "localhost"))

	c.PostServiceNewPort = cast.ToInt(getOrReturnDefault("POST_SERVICE_NEW_PORT", 9101))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
