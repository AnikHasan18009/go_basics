package main

import (
	"log"
	configuration "user-service/config"
	database "user-service/db"
	"user-service/web"
)

var config configuration.Config

func main() {
	var err error
	if config, err = configuration.LoadConfig("../config.json"); err != nil {
		log.Fatal("error loading config", err)
	}
	if err = database.EstablishConnection(config); err != nil {
		log.Fatal("error connecting database", err)
	}
	if err = web.StartServer(config.PORT); err != nil {
		log.Fatal("error starting server", err)
	}
	defer database.CloseConnection()

}
