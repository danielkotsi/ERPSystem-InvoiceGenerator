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
	GetInvoiceInfo(ctx context.Context, invoicetype string) (invoiceinfo models.InvoiceHTMLinfo, err error)
	CompleteInvoice(ctx context.Context, invo *models.Invoice) error
	UpdateDB(ctx context.Context, buyerNewBalance float64, buyerCodeNumber, invoicetype, aa string) error
	MakePDF(ctx context.Context, invo *models.Invoice) (pdf []byte, err error)
}
