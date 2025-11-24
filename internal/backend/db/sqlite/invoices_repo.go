package sqlite

import (
	"bytes"
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/utils"
	"fmt"
	"html/template"
	"log"
)

type InvoiceRepo struct {
	DB *sql.DB
}

func NewInvoiceRepo(db *sql.DB) *InvoiceRepo {
	return &InvoiceRepo{DB: db}
}

func (r *InvoiceRepo) DesignInvoice() error {

	tmpl, err := template.ParseFiles("../../assets/templates/home.page.html")
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		log.Println(err)
	}
	utils.HTMLtoPDF(buf.String(), "../../../../config/thenewpdffile.pdf")

	return nil
}

func (r *InvoiceRepo) ListProducts(ctx context.Context, search string) ([]models.Product, error) {
	search = fmt.Sprintf("%v%%", search)
	fmt.Println(search)
	query := "SELECT products.name,products.description,products.sku,products.unit_price,product_categories.name from products join product_categories on category_id==product_categories.id where sku LIKE ? and active==1"

	rows, err := r.DB.QueryContext(ctx, query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.Name, &p.Description, &p.Product_code, &p.Unit_price, &p.Category); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, nil
}

func (r *InvoiceRepo) ListCustomers(ctx context.Context, search string) (models.Customers, error) {
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
