package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
	"fmt"
	"net/http"
)

type ProductsService struct {
	Products repository.Products_repo
}

func NewProductsService(in repository.Products_repo) *ProductsService {
	return &ProductsService{Products: in}
}

func (s *ProductsService) ListProducts(ctx context.Context, r *http.Request) (resp models.Products, err error) {

	search := r.URL.Query().Get("search")
	fmt.Println("hello this is the search", search)
	products, err := s.Products.ListProducts(ctx, search)
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}
