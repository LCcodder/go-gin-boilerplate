package config

import (
	"fmt"
	"os"
	"time"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
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
			host, port, user, password, dbname),
		JWTSecret:             "dasmkdasmkda",
		JWTExpiration:         time.Hour * 24 * 90,
		RedisConnectionString: "localhost:6379",
		RedisPassword:         "",
	}
}
