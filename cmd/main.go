package main

import (
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
