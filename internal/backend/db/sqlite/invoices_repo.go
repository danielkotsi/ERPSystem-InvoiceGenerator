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

func (r *InvoiceRepo) ListCustomers(ctx context.Context, search string) ([]models.Customer, error) {
	return []models.Customer{}, nil
}
