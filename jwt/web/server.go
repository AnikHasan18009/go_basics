package web

import (
	"fmt"
	"net/http"
	"user-service/web/middlewares"
)

func StartServer(port string) error {
	mux := http.NewServeMux()
	manager := middlewares.NewManager()
	InitializeRoutes(mux, manager)
	err := http.ListenAndServe(":"+port, mux)
	fmt.Println("Server Started")

	return err
}
