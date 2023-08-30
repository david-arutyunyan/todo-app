package repository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github/todo-app"
	"log"
	"strings"
)

type UsersSegmentsPostgres struct {
	db *sqlx.DB
}

func NewUsersSegmentsPostgres(db *sqlx.DB) *UsersSegmentsPostgres {
	return &UsersSegmentsPostgres{db: db}
}

func (r *UsersSegmentsPostgres) GetUserSegments(userId string) ([]todo.Segment, error) {
	var segments []todo.Segment

	query := fmt.Sprintf("SELECT s.id, s.name FROM %s us INNER JOIN %s s on us.segment_id = s.id WHERE us.user_id = $1",
		usersSegmentsTable, segmentsTable)
	err := r.db.Select(&segments, query, userId)

	return segments, err
}

func (r *UsersSegmentsPostgres) UpdateUserSegments(a todo.A) error {
	var segmentsIdToAdd []string
	var segmentsIdToDelete []string

	var segmentsNamesToAdd []string
	var segmentsNamesToDelete []string

	for _, v := range a.Add {
		segmentsNamesToAdd = append(segmentsNamesToAdd, v.Name)
	}

	for _, v := range a.Delete {
		segmentsNamesToDelete = append(segmentsNamesToDelete, v.Name)
	}

	query, args, err := sqlx.In(fmt.Sprintf("SELECT id FROM %s WHERE name IN (?)", segmentsTable), segmentsNamesToAdd)
	if err != nil {
		log.Fatal(err)
	}
	err = r.db.Select(&segmentsIdToAdd, r.db.Rebind(query), args...)

	query, args, err = sqlx.In(fmt.Sprintf("SELECT id FROM %s WHERE name IN (?)", segmentsTable), segmentsNamesToDelete)
	if err != nil {
		log.Fatal(err)
	}
	err = r.db.Select(&segmentsIdToDelete, r.db.Rebind(query), args...)

	vals1 := []interface{}{}
	query = fmt.Sprintf("INSERT INTO %s (id, user_id, segment_id) VALUES", usersSegmentsTable)

	for _, v := range segmentsIdToAdd {
		query += "(?, ?, ?),"
		vals1 = append(vals1, uuid.New().String(), a.Id, v)
	}
	query = strings.TrimSuffix(query, ",")

	_, err = r.db.Exec(r.db.Rebind(query), vals1...)

	query, args, err = sqlx.In(fmt.Sprintf("DELETE FROM %s WHERE user_id=%d AND segment_id IN (?)", usersSegmentsTable, a.Id), segmentsIdToDelete)

	_, err = r.db.Exec(r.db.Rebind(query), args...)

	return err
}
