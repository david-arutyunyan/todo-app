package service

import (
	"github/todo-app"
	"github/todo-app/pkg/repository"
)

type User interface {
	CreateUser(user todo.User) (string, error)
	DeleteUser(userId string) error
}

type Segment interface {
	CreateSegment(segment todo.Segment) (string, error)
	DeleteSegment(segment todo.Segment) error
}

type UsersSegments interface {
	GetUserSegments(userId string) ([]todo.Segment, error)
	UpdateUserSegments(a todo.A) error
}

type Service struct {
	User
	Segment
	UsersSegments
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:          NewUserService(repo.User),
		Segment:       NewSegmentService(repo.Segment),
		UsersSegments: NewUsersSegmentsService(repo.UsersSegments),
	}
}
