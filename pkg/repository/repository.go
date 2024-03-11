package repository

import (
	"github.com/jmoiron/sqlx"
)

type Operation interface {
	GetData()
}

type Repository struct {
	Operation
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Operation: NewOperationPostgres(db),
	}
}
