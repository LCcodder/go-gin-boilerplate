package config

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "_"
	dbname   = "demo"
)

var (
	Config AppConfig = AppConfig{}
)

type AppConfig struct {
	PostgresConnectionString string
	JWTSecret                string
	JWTExpiration            string
}

func InitConfig() {
	Config = AppConfig{
		PostgresConnectionString: fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname),
		JWTSecret:     "dasmkdasmkda",
		JWTExpiration: "",
	}
}
