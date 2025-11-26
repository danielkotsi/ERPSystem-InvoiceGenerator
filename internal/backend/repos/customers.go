package repository

import (
	"context"
	"-invoice_manager/internal/backend/models"
)

type Customers_repo interface {
	ListCustomers(ctx context.Context, search string) (models.Customers, error)
	CreateCustomer(ctx context.Context, customer_data models.Customer) error
}
