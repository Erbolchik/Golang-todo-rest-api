package main

import (
	"log"

	todo "github.com/Erbolchik/Golang-todo-rest-api"
	"github.com/Erbolchik/Golang-todo-rest-api/pkg/handler"
	"github.com/Erbolchik/Golang-todo-rest-api/pkg/repository"
	"github.com/Erbolchik/Golang-todo-rest-api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
