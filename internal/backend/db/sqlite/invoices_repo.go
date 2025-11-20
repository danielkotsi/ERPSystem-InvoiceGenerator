package sqlite

import (
	"database/sql"
)

type InvoiceRepo struct {
	DB *sql.DB
}

func NewInvoiceRepo(db *sql.DB) *InvoiceRepo {
	return &InvoiceRepo{DB: db}
}
