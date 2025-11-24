package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
	"fmt"
	"net/http"
)

type InvoiceService struct {
	Invoice repository.Invoice_repo
}

func NewInvoiceService(in repository.Invoice_repo) *InvoiceService {
	return &InvoiceService{Invoice: in}
}

func (s *InvoiceService) ListCustomers(ctx context.Context, r *http.Request) (resp []models.Customer, err error) {

	search := r.URL.Query().Get("search")
	customers, err := s.Invoice.ListCustomers(ctx, search)
	if err != nil {
		return []models.Customer{}, err
	}
	return customers, nil
}

func (s *InvoiceService) ListProducts(ctx context.Context, r *http.Request) (resp models.Products, err error) {

	search := r.URL.Query().Get("search")
	fmt.Println("hello this is the search", search)
	products, err := s.Invoice.ListProducts(ctx, search)
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}
