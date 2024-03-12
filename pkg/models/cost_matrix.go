package models

type CostMatrix struct {
	Id              int `db:"id"`
	MicrocategoryId int `db:"microcategory_id"`
	LocationId      int `db:"location_name"`
	price           int `db:"price"`
}
