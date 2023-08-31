package service

import (
	_ "errors"
	"github/todo-app"
	"github/todo-app/pkg/repository"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user todo.User) (string, error) {
	return s.repo.CreateUser(user)

}

func (s *UserService) DeleteUser(userId string) error {
	return s.repo.DeleteUser(userId)

}
