package database

import (
	"database/sql"
	"fmt"

	"example.com/m/internal/config"
)

var Db *sql.DB

func ConnectToDatabase() {
	db, err := sql.Open("postgres", config.Config.PostgresConnectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	Db = db
	fmt.Println("Connected to database: " + config.Config.PostgresConnectionString)
}
