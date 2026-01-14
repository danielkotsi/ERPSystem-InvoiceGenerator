package sqlite

import (
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/utils"
	"errors"
	"fmt"
	"log"
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
	var productCat sql.NullString
	query := `SELECT 
	products.CodeNumber,
	products.name,
	products.description,
	products.unit_net_price,
	products.measurmentUnit,
	measurementUnits.unit,
	products.vat_category,
	categoriesforproducts.name 
	from products 
	left join product_categories on products.CodeNumber==product_categories.product_id 
	left join categoriesforproducts on product_categories.category_id==categoriesforproducts.id
	left join measurementUnits on products.measurmentUnit==measurementUnits.id 
	where products.name LIKE ? ;`

	rows, err := r.DB.QueryContext(ctx, query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.CodeNumber, &p.Name, &p.Description, &p.Unit_Net_Price, &p.MeasurementUnitCode, &p.MeasurementUnit, &p.VatCategory, &productCat); err != nil {
			return nil, err
		}
		p.ProductCategory = utils.NullableString(productCat)
		out = append(out, p)
	}
	return out, nil
}

func (r *ProductsRepo) CreateProduct(ctx context.Context, product_data models.Product) error {
	query := "insert into products(CodeNumber,name,description,unit_net_price,measurmentUnit,vat_category) values(?,?,?,?,?,?)"

	_, err := r.DB.ExecContext(ctx, query, product_data.CodeNumber, product_data.Name, product_data.Description, product_data.Unit_Net_Price, product_data.MeasurementUnitCode, product_data.VatCategory)
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
