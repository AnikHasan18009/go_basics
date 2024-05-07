package web

import (
	handler "user-service/web/handlers"
	"user-service/web/middlewares"

	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(handler.GetUsers),
		),
	)

	mux.Handle(
		"POST /user/create",
		manager.With(
			http.HandlerFunc(handler.CreateUser),
		),
	)
}
