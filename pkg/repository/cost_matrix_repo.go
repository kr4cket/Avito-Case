package repository

import "github.com/jmoiron/sqlx"

type CostMatrixPostgres struct {
	db *sqlx.DB
}

func NewCostMatrixPostgres(db *sqlx.DB) *CostMatrixPostgres {
	return &CostMatrixPostgres{db: db}
}

func (r *CostMatrixPostgres) GetData(userId int, locationId int, microcategoryId int) (map[string]string, error) {
	return map[string]string{
		"price":            "1000",
		"location_id":      "10",
		"microcategory_id": "5",
		"matrix_id":        "100",
		"user_segment_id":  "5",
	}, nil

}
