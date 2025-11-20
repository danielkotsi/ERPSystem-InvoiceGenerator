package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"

	// "-invoice_manager/internal/utils"
	// "encoding/json"
	// "log"
	"net/http"
	// "strings"
	// "time"
)

type InvoiceService struct {
	invoice repository.Invoice_repo
}

func NewInvoiceService(in repository.Invoice_repo) *InvoiceService {
	return &InvoiceService{invoice: in}
}

func (s *InvoiceService) List(ctx context.Context, r *http.Request) (resp models.Invoice, err error) {

	return models.Invoice{Something: "hello there this set up is working properly"}, nil
}
