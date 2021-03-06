package server

import (
	todo "github.com/Erbolchik/Golang-todo-rest-api"
	"github.com/Erbolchik/Golang-todo-rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, err)
}
type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		AuthService: newAuthService(repos.Authorization),
	}
}
