package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
	"-invoice_manager/internal/utils"
	"fmt"
	"net/http"
)

type CustomersService struct {
	Customers repository.Customers_repo
}

func NewCustomersService(in repository.Customers_repo) *CustomersService {
	return &CustomersService{Customers: in}
}

func (s *CustomersService) ListCustomers(ctx context.Context, r *http.Request) (resp []models.Customer, err error) {

	search := r.URL.Query().Get("search")
	customers, err := s.Customers.ListCustomers(ctx, search)
	if err != nil {
		return []models.Customer{}, err
	}
	return customers, nil
}

func (s *CustomersService) CreateCustomer(ctx context.Context, r *http.Request) error {

	var customer models.Customer
	if err := utils.ParseFormData(r, &customer); err != nil {
		return err
	}

	fmt.Println(customer)
	if err := s.Customers.CreateCustomer(r.Context(), customer); err != nil {
		return err
	}
	return nil
}
