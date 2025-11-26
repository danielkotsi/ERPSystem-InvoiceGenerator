package repository

import (
	"context"
	"-invoice_manager/internal/backend/models"
)

type Products_repo interface {
	ListProducts(ctx context.Context, search string) ([]models.Product, error)
}
