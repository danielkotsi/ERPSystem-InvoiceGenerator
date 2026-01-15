package product

import (
	"context"
	"-invoice_manager/internal/backend/product/models"
)

type Products_repo interface {
	ListProducts(ctx context.Context, search string) ([]models.Product, error)
	CreateProduct(ctx context.Context, product_data models.Product) error
}
