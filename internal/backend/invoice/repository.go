package invoice

import (
	"context"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/payload"
)

type Invoice_repo interface {
	GetInvoiceInfo(ctx context.Context, invoicetype string) (invoiceinfo models.InvoiceHTMLinfo, err error)
	CompleteInvoice(ctx context.Context, invo *payload.Invoice) error
	UpdateDB(ctx context.Context, buyerNewBalance float64, buyerCodeNumber, invoicetype, aa string) error
	MakePDF(ctx context.Context, invo *payload.Invoice) (pdf []byte, err error)
}

type Invoice_type interface {
	CalculateAlltheInvoiceLines(invoicetype string, paymentmethods *payload.PaymentMethods, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company) error
}
