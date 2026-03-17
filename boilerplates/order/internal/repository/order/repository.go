package order

import (
	"database/sql"
)

type repository struct{
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}