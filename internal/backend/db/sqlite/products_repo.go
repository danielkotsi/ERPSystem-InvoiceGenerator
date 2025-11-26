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
	query := "SELECT products.name,products.description,products.sku,products.unit_price,categoriesforproducts.name from products join product_categories on products.id==product_categories.product_id join categoriesforproducts on product_categories.category_id==categoriesforproducts.id where sku LIKE ? and active==1;"

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

func (r *ProductsRepo) CreateProduct(ctx context.Context, product_data models.Product) error {
	query := "insert into products(name,description,sku,unit_price,currency) values(?,?,?,?,?)"

	_, err := r.DB.ExecContext(ctx, query, product_data.Name, product_data.Description, product_data.Product_code, product_data.Unit_price, product_data.Currency)
	if err != nil {
		return err
	}
	return nil
}
