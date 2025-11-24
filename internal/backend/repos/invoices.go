package repository

import (
// "context"
// "-invoice_manager/internal/backend/models"
// "time"
)

type Invoice_repo interface {
	// ListCustomers(ctx context.Context, search string) ([]models.Customer, error)
	// ListProducts(ctx context.Context, search string) ([]models.Product, error)
	// ListInvoices(ctx context.Context, search string) ([]models.Invoice, error)
	// GetMAPK(ctx context.Context, search string) (int, error)
	// CreateMAPK_QRcode(ctx context.Context, search string) (int, error)
	DesignInvoice() error
}
