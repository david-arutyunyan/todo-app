package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github/todo-app"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) DeleteUser(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, userId)

	return err
}
