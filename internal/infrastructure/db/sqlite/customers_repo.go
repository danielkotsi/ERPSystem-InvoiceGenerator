package sqlite

import (
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/customer/models"
	"-invoice_manager/internal/backend/invoice/payload"
	"fmt"
)

type CustomersRepo struct {
	DB *sql.DB

	Stmts *CustomerStmts
}

func NewCustomersRepo(db *sql.DB) (*CustomersRepo, error) {
	customerStmts, err := NewCustomerStmts(db)
	if err != nil {
		return nil, err
	}
	return &CustomersRepo{DB: db, Stmts: customerStmts}, nil
}

func (r *CustomersRepo) ListCustomers(ctx context.Context, search string) ([]payload.Company, error) {
	search = fmt.Sprintf("%v%%", search)

	rows, err := r.Stmts.SearchByName.QueryContext(ctx, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []payload.Company
	for rows.Next() {
		var p payload.Company
		p.Address = &payload.AddressType{}
		if err := rows.Scan(
			&p.CodeNumber,
			&p.Name,
			&p.DOI,
			&p.GEMI,
			&p.Phone,
			&p.Mobile_Phone,
			&p.Email,
			&p.PostalAddress.Naming,
			&p.PostalAddress.Cellnumber,
			&p.PostalAddress.PostalCode,
			&p.PostalAddress.City,
			&p.Address.Street,
			&p.Address.Number,
			&p.Address.PostalCode,
			&p.Address.City,
			&p.VatNumber,
			&p.Country,
			&p.Branch,
			&p.OldBalance,
			&p.Discount,
		); err != nil {
			return nil, err
		}
		customers = append(customers, p)
	}
	return customers, nil
}

func (r *CustomersRepo) ListCustomerSuggestions(ctx context.Context, search string) ([]models.CustomerSuggestion, error) {
	search = fmt.Sprintf("%v%%", search)

	rows, err := r.Stmts.CustomerSuggestionsByName.QueryContext(ctx, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.CustomerSuggestion
	for rows.Next() {
		var p models.CustomerSuggestion
		if err := rows.Scan(
			&p.CodeNumber,
			&p.Name,
		); err != nil {
			return nil, err
		}
		customers = append(customers, p)
	}
	return customers, nil
}

func (r *CustomersRepo) GetCustomerById(ctx context.Context, code string) (customer payload.Company, err error) {
	var p payload.Company
	p.Address = &payload.AddressType{}

	err = r.Stmts.SearchById.
		QueryRowContext(ctx, code).
		Scan(
			&p.CodeNumber,
			&p.Name,
			&p.DOI,
			&p.GEMI,
			&p.Phone,
			&p.Mobile_Phone,
			&p.Email,
			&p.PostalAddress.Naming,
			&p.PostalAddress.Cellnumber,
			&p.PostalAddress.PostalCode,
			&p.PostalAddress.City,
			&p.Address.Street,
			&p.Address.Number,
			&p.Address.PostalCode,
			&p.Address.City,
			&p.VatNumber,
			&p.Country,
			&p.Branch,
			&p.OldBalance,
			&p.Discount,
		)
	if err != nil {
		return customer, err
	}
	return p, nil
}

func (r *CustomersRepo) CreateCustomer(ctx context.Context, customer_data payload.Company) error {

	_, err := r.Stmts.CreateCustomer.ExecContext(ctx,
		customer_data.CodeNumber,
		customer_data.Name,
		customer_data.DOI,
		customer_data.GEMI,
		customer_data.Phone,
		customer_data.Mobile_Phone,
		customer_data.Email,
		customer_data.PostalAddress.Naming,
		customer_data.PostalAddress.Cellnumber,
		customer_data.PostalAddress.PostalCode,
		customer_data.PostalAddress.City,
		customer_data.Address.Street,
		customer_data.Address.Number,
		customer_data.Address.PostalCode,
		customer_data.Address.City,
		customer_data.VatNumber,
		customer_data.Country,
		customer_data.Branch,
		customer_data.OldBalance,
		customer_data.Discount,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomersRepo) ListBranchCompanies(ctx context.Context, company, search string) ([]payload.BranchCompany, error) {
	search = fmt.Sprintf("%v%%", search)
	fmt.Println(search)

	rows, err := r.Stmts.SearchBranch.QueryContext(ctx, company, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branchcompanies []payload.BranchCompany
	for rows.Next() {
		var p payload.BranchCompany
		if err := rows.Scan(
			&p.BranchCode,
			&p.CompanyCode,
			&p.Name,
			&p.Phone,
			&p.Mobile_Phone,
			&p.Email,
			&p.Address.Street,
			&p.Address.Number,
			&p.Address.PostalCode,
			&p.Address.City,
			&p.Country,
			&p.OldBalance,
		); err != nil {
			return nil, err
		}
		branchcompanies = append(branchcompanies, p)
	}
	return branchcompanies, nil
}

func (r *CustomersRepo) ListBranchCompanySuggestions(ctx context.Context, company, search string) ([]models.BranchSuggestion, error) {
	search = fmt.Sprintf("%v%%", search)
	fmt.Println(search)

	rows, err := r.Stmts.BranchSuggestions.QueryContext(ctx, company, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branchcompanies []models.BranchSuggestion
	for rows.Next() {
		var p models.BranchSuggestion
		if err := rows.Scan(
			&p.CodeNumber,
			&p.CompanyCode,
			&p.Name,
		); err != nil {
			return nil, err
		}
		branchcompanies = append(branchcompanies, p)
	}
	return branchcompanies, nil
}
func (r *CustomersRepo) CreateBranchCompany(ctx context.Context, branch_data payload.BranchCompany) error {

	_, err := r.Stmts.CreateBranchCompany.ExecContext(ctx,
		branch_data.BranchCode,
		branch_data.CompanyCode,
		branch_data.Name,
		branch_data.Phone,
		branch_data.Mobile_Phone,
		branch_data.Email,
		branch_data.Address.Street,
		branch_data.Address.Number,
		branch_data.Address.PostalCode,
		branch_data.Address.City,
		branch_data.Country,
		0,
		branch_data.OldBalance,
		0,
	)
	if err != nil {
		return err
	}
	return nil
}
