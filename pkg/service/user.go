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

func (s *UserService) CreateUser(user todo.User) (int, error) {
	//user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

//func generatePasswordHash(password string) string {
//	hash := sha1.New()
//	hash.Write([]byte(password))
//
//	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
//}

func (s *UserService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}
