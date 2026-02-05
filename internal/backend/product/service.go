package product

import (
	"context"
	"invoice_manager/internal/backend/product/models"
)

type ProductsService struct {
	Products Products_repo
}

func NewProductsService(in Products_repo) *ProductsService {
	return &ProductsService{Products: in}
}

func (s *ProductsService) ListProducts(ctx context.Context, search string) (resp []models.Product, err error) {

	products, err := s.Products.ListProducts(ctx, search)
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func (s *ProductsService) CreateProduct(ctx context.Context, product models.Product) error {

	if err := s.Products.CreateProduct(ctx, product); err != nil {
		return err
	}
	return nil
}
