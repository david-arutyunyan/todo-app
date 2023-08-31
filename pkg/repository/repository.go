package repository

import (
	"github.com/jmoiron/sqlx"
	"github/todo-app"
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
	UpdateUserSegments(a todo.AlteredUserSegments) error
}

type Repository struct {
	User
	Segment
	UsersSegments
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:          NewUserPostgres(db),
		Segment:       NewSegmentPostgres(db),
		UsersSegments: NewUsersSegmentsPostgres(db),
	}
}
