package services

import (
	// "context"
	// "-invoice_manager/internal/backend/models"
	"context"
	"-invoice_manager/internal/backend/repos"
	"net/http"
)

type InvoiceService struct {
	Invoice repository.Invoice_repo
}

func NewInvoiceService(in repository.Invoice_repo) *InvoiceService {
	return &InvoiceService{Invoice: in}
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, r *http.Request) (pdf []byte, err error) {
	return pdf, nil
}
