package customer

import (
	"context"
	"invoice_manager/internal/backend/invoice/payload"
)

type Customers_repo interface {
	ListCustomers(ctx context.Context, search string) ([]payload.Company, error)
	GetCustomerById(ctx context.Context, code string) (payload.Company, error)
	ListBranchCompanies(ctx context.Context, company, search string) ([]payload.BranchCompany, error)
	CreateCustomer(ctx context.Context, customer_data payload.Company) error
	CreateBranchCompany(ctx context.Context, branch_data payload.BranchCompany) error
}
