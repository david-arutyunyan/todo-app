package repository

import (
	"github.com/jmoiron/sqlx"
	"github/todo-app"
)

type User interface {
	CreateUser(user todo.User) (int, error)
	DeleteUser(userId int) error
}

type Segment interface {
	CreateSegment(segment todo.Segment) (int, error)
	DeleteSegment(segment todo.Segment) error
}

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Segment: NewSegmentPostgres(db),
	}
}
