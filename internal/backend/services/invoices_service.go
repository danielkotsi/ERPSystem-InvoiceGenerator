package services

import (
	// "context"
	// "-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
)

type InvoiceService struct {
	Invoice repository.Invoice_repo
}

func NewInvoiceService(in repository.Invoice_repo) *InvoiceService {
	return &InvoiceService{Invoice: in}
}
