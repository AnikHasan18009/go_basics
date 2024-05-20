package db

import (
	"fmt"
	"log"
	configuration "user-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func EstablishConnection(config configuration.Config) error {
	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host) with your own
	args := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", config.HOST, config.DB_PORT, config.DB_NAME, config.DB_USER, config.DB_PASSWORD, config.SSL)
	var err error
	if Db, err = sqlx.Connect("postgres", args); err != nil {
		return err
	}

	if err = Db.Ping(); err != nil {
		return err
	}

	log.Println("Successfully Connected")
	return err

}
func CloseConnection() {
	Db.Close()
}
