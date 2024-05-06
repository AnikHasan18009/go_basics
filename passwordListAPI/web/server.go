package web

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	openRoutes()
	fmt.Println("Server Started")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
