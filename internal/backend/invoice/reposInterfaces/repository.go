package reposinterfaces

import (
	"context"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/backend/invoice/types"
)

type Invoice_repo interface {
	GetInvoiceInfo(ctx context.Context, invoicetype types.InvoiceType) (invoiceinfo models.InvoiceHTMLinfo, err error)
	HydrateInvoice(ctx context.Context, invo Invoice_type) error
	UpdateDB(ctx context.Context, buyerNewBalance float64, buyerCodeNumber, invoicetype, aa string) error
	Save(ctx context.Context, invo Invoice_type) error
}

type Invoice_type interface {
	CalculateInvoiceLines() error
	GetInvoice() (payload *payload.Invoice)
	MakePDF(ctx context.Context) (pdf []byte, err error)
}
