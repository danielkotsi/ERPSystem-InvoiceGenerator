package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
	"-invoice_manager/internal/utils"
	"net/http"
)

type CustomersService struct {
	Customers repository.Customers_repo
}

func NewCustomersService(in repository.Customers_repo) *CustomersService {
	return &CustomersService{Customers: in}
}

func (s *CustomersService) ListCustomers(ctx context.Context, r *http.Request) (resp []models.Company, err error) {

	search := r.URL.Query().Get("search")
	customers, err := s.Customers.ListCustomers(ctx, search)
	if err != nil {
		return []models.Company{}, err
	}
	return customers, nil
}

func (s *CustomersService) ListBranchCompanies(ctx context.Context, r *http.Request) (resp []models.BranchCompany, err error) {

	search := r.URL.Query().Get("search")
	company := r.URL.Query().Get("company")
	customers, err := s.Customers.ListBranchCompanies(ctx, company, search)
	if err != nil {
		return []models.BranchCompany{}, err
	}
	return customers, nil
}

func (s *CustomersService) CreateCustomer(ctx context.Context, r *http.Request) error {

	var customer models.Customer
	if err := utils.ParseFormData(r, &customer); err != nil {
		return err
	}

	if err := s.Customers.CreateCustomer(r.Context(), customer); err != nil {
		return err
	}
	return nil
}
