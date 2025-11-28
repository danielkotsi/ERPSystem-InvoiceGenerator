package sqlite

import (
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"errors"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
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
	product_id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	query := "insert into products(id,name,description,sku,unit_price,vat_category) values(?,?,?,?,?,?)"

	_, err = r.DB.ExecContext(ctx, query, product_id, product_data.Name, product_data.Description, product_data.Product_code, product_data.Unit_price, product_data.Vat_Category)
	if err != nil {
		return err
	}
	err = r.InsertProductIntoCategories(ctx, product_id.String(), product_data.Category)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductsRepo) InsertProductIntoCategories(ctx context.Context, product_id string, product_category string) error {
	err, id := r.GetProductCategoryID(ctx, product_category)
	if err != nil {
		return err
	}

	query := "insert into product_categories(product_id,category_id)values(?,?)"
	_, err = r.DB.ExecContext(ctx, query, product_id, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductsRepo) GetProductCategoryID(ctx context.Context, product_category string) (err error, id string) {
	query := "select id from categoriesforproducts where name==?"

	rows, err := r.DB.QueryContext(ctx, query, product_category)
	if err != nil {
		log.Println(err)
		return errors.New("Category For Product does not exist"), ""
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return err, ""
		}
	}
	return nil, id
}
