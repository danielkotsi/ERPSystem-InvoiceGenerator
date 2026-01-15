package product

import (
	"context"
	"-invoice_manager/internal/backend/product/models"
	"-invoice_manager/internal/utils"
	"fmt"
	"net/http"
)

type ProductsService struct {
	Products Products_repo
}

func NewProductsService(in Products_repo) *ProductsService {
	return &ProductsService{Products: in}
}

func (s *ProductsService) ListProducts(ctx context.Context, r *http.Request) (resp []models.Product, err error) {

	search := r.URL.Query().Get("search")
	fmt.Println("hello this is the search", search)
	products, err := s.Products.ListProducts(ctx, search)
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (s *ProductsService) CreateProduct(ctx context.Context, r *http.Request) error {

	var product models.Product
	if err := utils.ParseFormData(r, &product); err != nil {
		return err
	}

	fmt.Println("hello this is the product", product)
	if err := s.Products.CreateProduct(r.Context(), product); err != nil {
		return err
	}
	return nil
}
