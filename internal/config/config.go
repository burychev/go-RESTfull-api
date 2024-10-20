package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		AppConfig      AppConfig
		PostgresConfig PostgresConfig
	}

	AppConfig struct {
		Host string
		Port string
	}

	PostgresConfig struct {
		Database string
		User     string
		Password string
	}
)

func GetConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	c := &Config{}
	c.AppConfig.Host = os.Getenv("HOST")
	c.AppConfig.Port = os.Getenv("PORT")

	// PostgresConfig
	c.PostgresConfig.Database = os.Getenv("PG_DBNAME")
	c.PostgresConfig.User = os.Getenv("PG_USERNAME")
	c.PostgresConfig.Password = os.Getenv("PG_PASSWORD")

	return c
}
