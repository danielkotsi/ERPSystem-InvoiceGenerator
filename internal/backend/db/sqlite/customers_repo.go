package sqlite

import (
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"fmt"
)

type CustomersRepo struct {
	DB *sql.DB
}

func NewCustomersRepo(db *sql.DB) *CustomersRepo {
	return &CustomersRepo{DB: db}
}

func (r *CustomersRepo) ListCustomers(ctx context.Context, search string) (models.Customers, error) {
	search = fmt.Sprintf("%v%%", search)
	fmt.Println(search)
	query := "SELECT name,address_line1,address_num1,address_line2,address_num2,city,state,postal_code,country,email,phone,mobile_phone,tax_id from companies  where name LIKE ? "

	rows, err := r.DB.QueryContext(ctx, query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers models.Customers
	for rows.Next() {
		var p models.Customer
		if err := rows.Scan(&p.Name, &p.Address1, &p.NumofAdd1, &p.Address2, &p.NumofAdd2, &p.City, &p.State, &p.Postal_code, &p.Country, &p.Email, &p.Phone, &p.Mobile_Phone, &p.VAT); err != nil {
			return nil, err
		}
		customers = append(customers, p)
	}
	return customers, nil
}

func (r *CustomersRepo) CreateCustomer(ctx context.Context, customer_data models.Customer) error {
	return nil
}
