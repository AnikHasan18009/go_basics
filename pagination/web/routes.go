package web

import (
	handler "user-service/web/handlers"
	"user-service/web/middlewares"

	"net/http"
)

func InitializeRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(handler.GetUsers),
		),
	)

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handler.CreateNewUser),
		),
	)

	mux.Handle(
		"GET /users/{id}",
		manager.With(
			http.HandlerFunc(handler.GetRequestedUser),
		),
	)
	mux.Handle(
		"DELETE /users/{id}",
		manager.With(
			http.HandlerFunc(handler.RemoveRequestedUser),
		),
	)
	mux.Handle(
		"PATCH /users/{id}",
		manager.With(
			http.HandlerFunc(handler.UpdateRequestedUser),
		),
	)
}
