package repo

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository interface {
}
type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) (Repository, error) {
	return &repo{
		db: db,
	}, nil
}