package sqlite

import (
	"context"
	// "crypto/rand"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	// "-invoice_manager/internal/utils"
	"fmt"
	// "math/big"
)

type CustomersRepo struct {
	DB *sql.DB
}

func NewCustomersRepo(db *sql.DB) *CustomersRepo {
	return &CustomersRepo{DB: db}
}

func (r *CustomersRepo) ListCustomers(ctx context.Context, search string) ([]models.Company, error) {
	search = fmt.Sprintf("%v%%", search)
	fmt.Println(search)
	query := "SELECT CodeNumber, NAME,DOI,GEMI,Phone,Mobile_Phone,Email,PostalCellName,PostalCellNumber,PostalCellPostalCode,PostalCellCity,AddStreet,AddNumber, AddPostalCode,AddCity,VatNumber,Country,Branch,Balance,Discount from customers  where NAME LIKE ? "

	rows, err := r.DB.QueryContext(ctx, query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Company
	for rows.Next() {
		var p models.Company
		if err := rows.Scan(&p.CodeNumber, &p.Name, &p.DOI, &p.GEMI, &p.Phone, &p.Mobile_Phone, &p.Email, &p.PostalAddress.Naming, &p.PostalAddress.Cellnumber, &p.PostalAddress.PostalCode, &p.PostalAddress.City, &p.Address.Street, &p.Address.Number, &p.Address.PostalCode, &p.Address.City, &p.VatNumber, &p.Country, &p.Branch, &p.OldBalance, &p.Discount); err != nil {
			return nil, err
		}
		customers = append(customers, p)
	}
	fmt.Println(customers)
	return customers, nil
}

func (r *CustomersRepo) CreateCustomer(ctx context.Context, customer_data models.Customer) error {
	// x := rand.Reader
	// y, _ := rand.Int(x, big.NewInt(2000))
	// code := fmt.Sprintf("%s-%s", customer_data.Name[0:3], y.String())

	query := "insert into companies(name,address_street,address_number,city,postal_code,country,email,phone,mobile_phone,vat_number,branch,entity_type) values(?,?,?,?,?,?,?,?,?,?,?,?) "

	_, err := r.DB.ExecContext(ctx, query, customer_data.Name, customer_data.Address.Street, customer_data.Address.Number, customer_data.Address.City, customer_data.Address.PostalCode, customer_data.Country, customer_data.Email, customer_data.Phone, customer_data.Mobile_Phone, customer_data.VatNumber, customer_data.Branch, customer_data.EntityType)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomersRepo) ListBranchCompanies(ctx context.Context, company, search string) ([]models.BranchCompany, error) {
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

	var branchcompanies []models.BranchCompany
	for rows.Next() {
		var p models.BranchCompany
		if err := rows.Scan(&p.BranchCode, &p.CompanyCode, &p.Name, &p.Phone, &p.Mobile_Phone, &p.Email, &p.Address.Street, &p.Address.Number, &p.Address.PostalCode, &p.Address.City, &p.Country, &p.OldBalance); err != nil {
			return nil, err
		}
		p.Name = fmt.Sprintf("%s %s", p.BranchCode, p.Name)
		branchcompanies = append(branchcompanies, p)
	}
	fmt.Println(branchcompanies)
	return branchcompanies, nil
}
