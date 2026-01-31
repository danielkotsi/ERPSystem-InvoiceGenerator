package customer

import (
	"context"
	"-invoice_manager/internal/backend/customer/models"
	"-invoice_manager/internal/backend/invoice/payload"
)

type Customers_repo interface {
	ListCustomers(ctx context.Context, search string) ([]payload.Company, error)
	ListCustomerSuggestions(ctx context.Context, search string) ([]models.CustomerSuggestion, error)
	GetCustomerById(ctx context.Context, code string) (payload.Company, error)
	ListBranchCompanies(ctx context.Context, company, search string) ([]payload.BranchCompany, error)
	ListFullBranchCompany(ctx context.Context, company, search string) (payload.BranchCompany, error)
	ListBranchCompanySuggestions(ctx context.Context, company, search string) ([]models.BranchSuggestion, error)
	CreateCustomer(ctx context.Context, customer_data payload.Company) error
	CreateBranchCompany(ctx context.Context, branch_data payload.BranchCompany) error
}
