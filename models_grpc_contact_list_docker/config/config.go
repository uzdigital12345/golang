package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	EventServiceHost string
	EventServicePort    int
}

func Load() (Config) {
	c := Config{}

	c.PostgresHost = cast.ToString( getOrReturnDefault("POSTGRES_HOST", "127.0.0.1"))
  	c.PostgresPort = cast.ToInt( getOrReturnDefault("POSTGRES_PORT", 9898) )
  	c.PostgresDatabase = cast.ToString( getOrReturnDefault("POSTGRES_DATABASE", "contact_list_info"))
  	c.PostgresUser = cast.ToString( getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString( getOrReturnDefault("POSTGRES_PASSWORD", "123"))
	c.EventServiceHost = cast.ToString(getOrReturnDefault("EVENT_SERVICE_HOST", "127.0.0.1"))
	c.EventServicePort = cast.ToInt(getOrReturnDefault("EVENT_SERVICE_PORT", 9090))	  

	return c
}

func getOrReturnDefault(key string, default_value interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return default_value
}