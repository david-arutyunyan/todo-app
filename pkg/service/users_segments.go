package service

import (
	"github/todo-app"
	"github/todo-app/pkg/repository"
)

type UsersSegmentsService struct {
	repo repository.UsersSegments
}

func NewUsersSegmentsService(repo repository.UsersSegments) *UsersSegmentsService {
	return &UsersSegmentsService{repo: repo}
}

func (s *UsersSegmentsService) GetUserSegments(userId string) ([]todo.Segment, error) {
	return s.repo.GetUserSegments(userId)
}

func (s *UsersSegmentsService) UpdateUserSegments(a todo.AlteredUserSegments) error {
	return s.repo.UpdateUserSegments(a)
}
