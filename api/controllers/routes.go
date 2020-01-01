package controllers

import "github.com/alcjohn/api_go/api/middlewares"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")
}
