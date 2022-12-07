package service

import "github.com/NurzhanIT/Go_todo_api.git/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
