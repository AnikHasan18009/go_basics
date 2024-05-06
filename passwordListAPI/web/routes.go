package web

import (
	handler "passwordList/web/handlers"

	"net/http"
)

func openRoutes() {

	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/list", handler.ReturnList)

}
