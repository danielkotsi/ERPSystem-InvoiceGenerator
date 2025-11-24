package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"

	// "-invoice_manager/internal/utils"
	// "encoding/json"
	// "log"
	"html/template"
	"net/http"
	// "strings"
	// "time"
)

type InvoiceService struct {
	Invoice repository.Invoice_repo
	Tmpl    *template.Template
}

func NewInvoiceService(in repository.Invoice_repo, tmpl *template.Template) *InvoiceService {
	return &InvoiceService{Invoice: in, Tmpl: tmpl}
}

func (s *InvoiceService) ListCustomers(ctx context.Context, r *http.Request) (resp models.Invoice, err error) {

	return models.Invoice{}, nil
}
