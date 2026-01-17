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

func (s *CustomersService) ListCustomers(ctx context.Context, search string) (resp []payload.Company, err error) {
	customers, err := s.Customers.ListCustomers(ctx, search)
	if err != nil {
		return []payload.Company{}, err
	}
	return customers, nil
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

func (s *CustomersService) ListBranchCompanies(ctx context.Context, search, company string) (resp []payload.BranchCompany, err error) {

	customers, err := s.Customers.ListBranchCompanies(ctx, company, search)
	if err != nil {
		return []payload.BranchCompany{}, err
	}
	return customers, nil
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
