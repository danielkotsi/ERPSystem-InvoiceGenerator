package customer

import (
	"context"
	"-invoice_manager/internal/backend/customer/models"
	"-invoice_manager/internal/backend/invoice/payload"
	"fmt"
)

type CustomersService struct {
	Customers Customers_repo
}

func NewCustomersService(in Customers_repo) *CustomersService {
	return &CustomersService{Customers: in}
}

// returns full info for Customers for the Customers Page
func (s *CustomersService) ListCustomers(ctx context.Context, search string) (resp []payload.Company, err error) {
	customers, err := s.Customers.ListCustomers(ctx, search)
	if err != nil {
		return []payload.Company{}, err
	}
	return customers, nil
}

// returns basic Company Info just for the suggestions
func (s *CustomersService) ListCustomerSuggestions(ctx context.Context, search string) (resp []models.CustomerSuggestion, err error) {
	customers, err := s.Customers.ListCustomerSuggestions(ctx, search)
	if err != nil {
		return []models.CustomerSuggestion{}, err
	}
	return customers, nil
}

// returns Complete Branch Company Info
func (s *CustomersService) ListBranchCompanies(ctx context.Context, search, company string) (resp []payload.BranchCompany, err error) {
	customers, err := s.Customers.ListBranchCompanies(ctx, company, search)
	if err != nil {
		return []payload.BranchCompany{}, err
	}
	return customers, nil
}
func (s *CustomersService) GetBranchCompanyById(ctx context.Context, search, company string) (resp payload.BranchCompany, err error) {
	branchCompany, err := s.Customers.ListFullBranchCompany(ctx, company, search)
	if err != nil {
		return resp, err
	}
	return branchCompany, nil
}

// returns basic Branch Company Info just for the suggestions
func (s *CustomersService) ListBranchCompanySuggestions(ctx context.Context, search, company string) (resp []models.BranchSuggestion, err error) {
	customers, err := s.Customers.ListBranchCompanySuggestions(ctx, company, search)
	if err != nil {
		return []models.BranchSuggestion{}, err
	}
	return customers, nil
}

func (s *CustomersService) GetCustomerByIdWithoutBranchCompanies(ctx context.Context, codeNumber string) (resp payload.Company, err error) {
	customer, err := s.Customers.GetCustomerById(ctx, codeNumber)
	if err != nil {
		return resp, err
	}
	return customer, nil
}
func (s *CustomersService) GetCustomerById(ctx context.Context, codeNumber string) (resp models.CustomerById, err error) {
	customer, err := s.Customers.GetCustomerById(ctx, codeNumber)
	if err != nil {
		return resp, err
	}
	branchCompanies, err := s.Customers.ListBranchCompanies(ctx, codeNumber, "")
	if err != nil {
		return resp, err
	}
	resp.Customer = customer
	resp.BranchCompanies = branchCompanies
	fmt.Println(resp)
	return resp, nil
}

func (s *CustomersService) CreateCustomer(ctx context.Context, customer payload.Company) error {
	if err := s.Customers.CreateCustomer(ctx, customer); err != nil {
		return err
	}
	return nil
}

func (s *CustomersService) CreateBranchCompany(ctx context.Context, branch payload.BranchCompany) (err error) {
	if err := s.Customers.CreateBranchCompany(ctx, branch); err != nil {
		return err
	}
	return nil
}
