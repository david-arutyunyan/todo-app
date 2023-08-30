package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github/todo-app"
)

type SegmentPostgres struct {
	db *sqlx.DB
}

func NewSegmentPostgres(db *sqlx.DB) *SegmentPostgres {
	return &SegmentPostgres{db: db}
}

func (r *SegmentPostgres) CreateSegment(segment todo.Segment) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (id, name) VALUES ($1, $2) RETURNING id", segmentsTable)

	row := r.db.QueryRow(query, uuid.New().String(), segment.Name)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *SegmentPostgres) DeleteSegment(segment todo.Segment) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE name=$1", segmentsTable)
	_, err := r.db.Exec(query, segment.Name)

	return err
}
