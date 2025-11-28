package sqlite

import (
	"context"
	"errors"
	// "crypto/rand"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"fmt"
	// "math/big"
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
	query := "SELECT name,address_street,address_number,city,postal_code,country,entity_type,branch,vat_number,email,phone,mobile_phone from companies  where name LIKE ? "

	rows, err := r.DB.QueryContext(ctx, query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers models.Customers
	for rows.Next() {
		var p models.Customer
		if err := rows.Scan(&p.Name, &p.Address.Street, &p.Address.Number, &p.Address.City, &p.Address.PostalCode, &p.Country, &p.EntityType, &p.Branch, &p.VatNumber, p.Email, p.Phone, &p.Mobile_Phone); err != nil {
			return nil, err
		}
		customers = append(customers, p)
	}
	fmt.Println(customers)
	return customers, errors.New("this is a test")
}

func (r *CustomersRepo) CreateCustomer(ctx context.Context, customer_data models.Customer) error {
	// x := rand.Reader
	// y, _ := rand.Int(x, big.NewInt(2000))
	// code := fmt.Sprintf("%s-%s", customer_data.Name[0:3], y.String())
	//
	// query := "insert into companies(code,name,address_line1,address_num1,address_line2,address_num2,city,state,postal_code,country,email,phone,mobile_phone,tax_id) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?) "
	//
	// _, err := r.DB.ExecContext(ctx, query, code, customer_data.Name, customer_data.Address1, customer_data.NumofAdd1, customer_data.Address2, customer_data.NumofAdd2, customer_data.City, customer_data.State, customer_data.Postal_code, customer_data.Country, customer_data.Email, customer_data.Phone, customer_data.Mobile_Phone, customer_data.VAT)
	// if err != nil {
	// 	return err
	// }
	return nil
}
