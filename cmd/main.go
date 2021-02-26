package main

import (
	"log"

	todo "github.com/Erbolchik/Golang-todo-rest-api"
	"github.com/Erbolchik/Golang-todo-rest-api/pkg/handler"
	"github.com/Erbolchik/Golang-todo-rest-api/pkg/repository"
	"github.com/Erbolchik/Golang-todo-rest-api/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host: "localhost",
		Port:"5436",
		Username:"postgres",
		Password:"qwerty",
		DBName:"postgres",
		SSLMode:"disable"
	})

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
