package web

import (
	"fmt"
	"log"
	"net/http"
	"user-service/web/middlewares"
)

func StartServer() {
	mux := http.NewServeMux()

	manager := middlewares.NewManager()
	InitRoutes(mux, manager)

	fmt.Println("Server Started")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
