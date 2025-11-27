package services

import (
	// "context"
	// "-invoice_manager/internal/backend/models"
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
	"-invoice_manager/internal/utils"
	"net/http"
)

type InvoiceService struct {
	Invoice repository.Invoice_repo
}

func NewInvoiceService(in repository.Invoice_repo) *InvoiceService {
	return &InvoiceService{Invoice: in}
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, r *http.Request) (pdf []byte, err error) {
	var invo models.Invoice
	err = utils.ParseFormData(r, &invo)
	if err != nil {
		return nil, err
	}
	pdf, err = s.Invoice.DesignInvoice(ctx, invo)
	if err != nil {
		return pdf, err
	}
	return pdf, nil
}
