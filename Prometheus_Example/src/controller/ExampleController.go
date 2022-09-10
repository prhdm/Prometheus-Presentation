package controller

import (
	database "PrometheusExample/src/database"
	"PrometheusExample/src/service"
)

type ExampleController struct {
	http service.HttpServer
}

var DB database.Database

func NewExampleController(http service.HttpServer) *ExampleController {
	DB = *database.NewDatabase()
	DB.CreateTable("users")
	DB.CreateTable("games")
	http.Register("/signup", signupHandler)
	http.Register("/login", loginHandler)
	http.Register("/logout", logoutHandler)
	http.Register("/createGame", createGameHandler)
	http.Register("/joinGame", joinGameHandler)
	http.Register("/exitGame", exitGameHandler)
	return &ExampleController{http: http}
}
