package service

import (
	"github/todo-app"
	"github/todo-app/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment todo.Segment) (string, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentService) DeleteSegment(segment todo.Segment) error {
	return s.repo.DeleteSegment(segment)
}
