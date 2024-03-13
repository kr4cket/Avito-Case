package repository

import "github.com/jmoiron/sqlx"

type CostMatrix interface {
	GetData(userId int, locationId int, microcategoryId int) (map[string]string, error)
}

type Repository struct {
	CostMatrix
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CostMatrix: NewCostMatrixPostgres(db),
	}
}
