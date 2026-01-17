package sqlite

import (
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/invoice/payload"
	"fmt"
)

type CustomersRepo struct {
	DB *sql.DB
}

func NewCustomersRepo(db *sql.DB) *CustomersRepo {
	return &CustomersRepo{DB: db}
}

func (r *CustomersRepo) ListCustomers(ctx context.Context, search string) ([]payload.Company, error) {
	search = fmt.Sprintf("%v%%", search)
	fmt.Println(search)
	query := "SELECT CodeNumber, NAME,DOI,GEMI,Phone,Mobile_Phone,Email,PostalCellName,PostalCellNumber,PostalCellPostalCode,PostalCellCity,AddStreet,AddNumber, AddPostalCode,AddCity,VatNumber,Country,Branch,Balance,Discount from customers  where NAME LIKE ? "

	rows, err := r.DB.QueryContext(ctx, query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []payload.Company
	for rows.Next() {
		var p payload.Company
		p.Address = &payload.AddressType{}
		if err := rows.Scan(&p.CodeNumber, &p.Name, &p.DOI, &p.GEMI, &p.Phone, &p.Mobile_Phone, &p.Email, &p.PostalAddress.Naming, &p.PostalAddress.Cellnumber, &p.PostalAddress.PostalCode, &p.PostalAddress.City, &p.Address.Street, &p.Address.Number, &p.Address.PostalCode, &p.Address.City, &p.VatNumber, &p.Country, &p.Branch, &p.OldBalance, &p.Discount); err != nil {
			return nil, err
		}
		customers = append(customers, p)
	}
	fmt.Println(customers)
	return customers, nil
}

func (r *CustomersRepo) GetCustomerById(ctx context.Context, code string) (customer payload.Company, err error) {
	query := "SELECT CodeNumber, NAME,DOI,GEMI,Phone,Mobile_Phone,Email,PostalCellName,PostalCellNumber,PostalCellPostalCode,PostalCellCity,AddStreet,AddNumber, AddPostalCode,AddCity,VatNumber,Country,Branch,Balance,Discount from customers  where CodeNumber== ? "

	rows, err := r.DB.QueryContext(ctx, query, code)
	if err != nil {
		return customer, err
	}
	defer rows.Close()

	var p payload.Company
	p.Address = &payload.AddressType{}
	for rows.Next() {
		if err := rows.Scan(&p.CodeNumber, &p.Name, &p.DOI, &p.GEMI, &p.Phone, &p.Mobile_Phone, &p.Email, &p.PostalAddress.Naming, &p.PostalAddress.Cellnumber, &p.PostalAddress.PostalCode, &p.PostalAddress.City, &p.Address.Street, &p.Address.Number, &p.Address.PostalCode, &p.Address.City, &p.VatNumber, &p.Country, &p.Branch, &p.OldBalance, &p.Discount); err != nil {
			return customer, err
		}
	}
	return p, nil
}

func (r *CustomersRepo) CreateCustomer(ctx context.Context, customer_data payload.Company) error {

	query := `insert into customers(
	CodeNumber,
	NAME,
	DOI,
	GEMI,
	Phone,
	Mobile_Phone,
	Email,
	PostalCellName  ,
	PostalCellNumber,
	PostalCellPostalCode  ,
	PostalCellCity  ,
	AddStreet  ,
	AddNumber  ,
	AddPostalCode  ,
	AddCity,
	VatNumber,
	Country,
	Branch,
	Balance,
	Discount
	) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) `

	_, err := r.DB.ExecContext(ctx, query, customer_data.CodeNumber, customer_data.Name, customer_data.DOI, customer_data.GEMI, customer_data.Phone, customer_data.Mobile_Phone, customer_data.Email, customer_data.PostalAddress.Naming, customer_data.PostalAddress.Cellnumber, customer_data.PostalAddress.PostalCode, customer_data.PostalAddress.City, customer_data.Address.Street, customer_data.Address.Number, customer_data.Address.PostalCode, customer_data.Address.City, customer_data.VatNumber, customer_data.Country, customer_data.Branch, customer_data.OldBalance, customer_data.Discount)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomersRepo) ListBranchCompanies(ctx context.Context, company, search string) ([]payload.BranchCompany, error) {
	search = fmt.Sprintf("%v%%", search)
	fmt.Println(search)
	query := `SELECT 
	BranchCode,
	CompanyCode,
	NAME,
	Phone,
	Mobile_Phone,
	Email,
	AddStreet,
	AddNumber,
	AddPostalCode,
	AddCity,
	Country,
	Balance
	from BranchCompanies 
	where CompanyCode==? and BranchCode like ?;`

	rows, err := r.DB.QueryContext(ctx, query, company, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var branchcompanies []payload.BranchCompany
	for rows.Next() {
		var p payload.BranchCompany
		if err := rows.Scan(&p.BranchCode, &p.CompanyCode, &p.Name, &p.Phone, &p.Mobile_Phone, &p.Email, &p.Address.Street, &p.Address.Number, &p.Address.PostalCode, &p.Address.City, &p.Country, &p.OldBalance); err != nil {
			return nil, err
		}
		branchcompanies = append(branchcompanies, p)
	}
	return branchcompanies, nil
}

func (r *CustomersRepo) CreateBranchCompany(ctx context.Context, branch_data payload.BranchCompany) error {

	query := `insert into BranchCompanies(
	BranchCode,
	CompanyCode,
	NAME,
	Phone,
	Mobile_Phone,
	Email,
	AddStreet  ,
	AddNumber  ,
	AddPostalCode  ,
	AddCity,
	Country,
	Branch,
	Balance,
	Discount
	) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?) `

	_, err := r.DB.ExecContext(ctx, query, branch_data.BranchCode, branch_data.CompanyCode, branch_data.Name, branch_data.Phone, branch_data.Mobile_Phone, branch_data.Email, branch_data.Address.Street, branch_data.Address.Number, branch_data.Address.PostalCode, branch_data.Address.City, branch_data.Country, 0, branch_data.OldBalance, 0)
	if err != nil {
		return err
	}
	return nil
}
