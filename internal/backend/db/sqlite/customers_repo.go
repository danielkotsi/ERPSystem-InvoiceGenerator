package sqlite

import (
	"context"
	// "crypto/rand"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/utils"
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
		var street, number, city, postlacode, email, phone, mobilephone sql.NullString
		if err := rows.Scan(&p.Name, &street, &number, &city, &postlacode, &p.Country, &p.EntityType, &p.Branch, &p.VatNumber, &email, &phone, &mobilephone); err != nil {
			return nil, err
		}

		if utils.CheckIfSomethingNotNull(street, number, city, postlacode) {
			p.Address = &models.AddressType{}
			p.Address.Street = utils.NullableString(street)
			p.Address.Number = utils.NullableString(number)
			p.Address.City = utils.NullableString(city)
			p.Address.PostalCode = utils.NullableString(postlacode)
		}
		p.Email = utils.NullableString(email)
		p.Phone = utils.NullableString(phone)
		p.Mobile_Phone = utils.NullableString(mobilephone)

		customers = append(customers, p)
	}
	fmt.Println(customers)
	return customers, nil
}

func (r *CustomersRepo) CreateCustomer(ctx context.Context, customer_data models.Customer) error {
	// x := rand.Reader
	// y, _ := rand.Int(x, big.NewInt(2000))
	// code := fmt.Sprintf("%s-%s", customer_data.Name[0:3], y.String())

	query := "insert into companies(name,address_street,address_number,city,postal_code,country,email,phone,mobile_phone,vat_number,branch,entity_type) values(?,?,?,?,?,?,?,?,?,?,?,?) "

	_, err := r.DB.ExecContext(ctx, query, customer_data.Name, customer_data.Address.Street, customer_data.Address.Number, customer_data.Address.City, customer_data.Address.PostalCode, customer_data.Country, customer_data.Email, customer_data.Phone, customer_data.Mobile_Phone, customer_data.VatNumber, customer_data.Branch, customer_data.EntityType)
	if err != nil {
		return err
	}
	return nil
}
