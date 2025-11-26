package sqlite

import (
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"fmt"
)

type ProductsRepo struct {
	DB *sql.DB
}

func NewProductsRepo(db *sql.DB) *ProductsRepo {
	return &ProductsRepo{DB: db}
}

func (r *ProductsRepo) ListProducts(ctx context.Context, search string) ([]models.Product, error) {
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
