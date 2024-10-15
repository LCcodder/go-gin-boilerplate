package config

import (
	"fmt"
	"os"
	"time"
)

var (
	Config AppConfig = AppConfig{}
)

type AppConfig struct {
	PostgresConnectionString string
	JWTSecret                string
	JWTExpiration            time.Duration
	RedisConnectionString    string
	RedisPassword            string
}

func InitConfig() {
	Config = AppConfig{
		PostgresConnectionString: fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME")),
		JWTSecret:             os.Getenv("JWT_SECRET"),
		JWTExpiration:         time.Hour * 24 * 90,
		RedisConnectionString: os.Getenv("REDIS_CONNECTION"),
		RedisPassword:         "",
	}
}
