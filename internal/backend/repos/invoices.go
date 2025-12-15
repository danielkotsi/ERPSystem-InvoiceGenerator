package repository

import (
	"context"
	"-invoice_manager/internal/backend/models"
)

// "context"
// "-invoice_manager/internal/backend/models"

type Invoice_repo interface {
	// ListInvoices(ctx context.Context, search string) ([]models.Invoice, error)
	// GetMAPK(ctx context.Context, search string) (int, error)
	// CreateMAPK_QRcode(ctx context.Context, search string) (int, error)
	CompleteInvoice(ctx context.Context, invo *models.Invoice) error
	MakePDF(ctx context.Context, invo *models.InvoicePayload) (pdf []byte, err error)
}
