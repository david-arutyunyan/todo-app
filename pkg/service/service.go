package service

import (
	"github/todo-app"
	"github/todo-app/pkg/repository"
)

type User interface {
	CreateUser(user todo.User) (int, error)
	DeleteUser(userId int) error
}

type Segment interface {
	CreateSegment(segment todo.Segment) (int, error)
	DeleteSegment(segment todo.Segment) error
}

type Service struct {
	User
	Segment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repo.User),
		Segment: NewSegmentService(repo.Segment),
	}
}
