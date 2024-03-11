package repository

import (
	"github.com/jmoiron/sqlx"
)

type OperationPostgres struct {
	db *sqlx.DB
}

func NewOperationPostgres(db *sqlx.DB) *OperationPostgres {
	return &OperationPostgres{db: db}
}

func (r *OperationPostgres) GetData() {
	//	TODO Получить инфу с БД
}
