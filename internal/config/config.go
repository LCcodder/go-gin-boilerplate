package config

import (
	"fmt"
	"time"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = ""
	dbname   = "demo"
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
