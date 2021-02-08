package main

import (
	"log"
	"github.com/Erbolchik/Golang-todo-rest-api"
	"github.com/Erbolchik/Golang-todo-rest-api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
