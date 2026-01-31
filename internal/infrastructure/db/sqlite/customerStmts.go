package sqlite

import (
	"database/sql"
)

type CustomerStmts struct {
	CreateCustomer            *sql.Stmt
	CreateBranchCompany       *sql.Stmt
	SearchByName              *sql.Stmt
	CustomerSuggestionsByName *sql.Stmt
	SearchBranch              *sql.Stmt
	SearchFullBranch          *sql.Stmt
	BranchSuggestions         *sql.Stmt
	SearchById                *sql.Stmt
}

func NewCustomerStmts(db *sql.DB) (*CustomerStmts, error) {
	searchFullBranchQuery := `
	SELECT 
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
	where CompanyCode==? and BranchCode==?;
`
	searchFullBranchStmt, err := db.Prepare(searchFullBranchQuery)
	if err != nil {
		return nil, err
	}
	branchSuggestionsQuery := `SELECT 
	BranchCode,
	CompanyCode,
	NAME
	from BranchCompanies 
	where CompanyCode==? and BranchCode like ?;
	`
	branchSuggestionsStmt, err := db.Prepare(branchSuggestionsQuery)
	if err != nil {
		return nil, err
	}
	customersSuggestionsByNameQuery := `SELECT CodeNumber,
	NAME
	from customers  where NAME LIKE ? 
	`
	customerSuggestionsStmt, err := db.Prepare(customersSuggestionsByNameQuery)
	if err != nil {
		return nil, err
	}
	createCustomerQuery := `insert into customers(
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
	insertStmt, err := db.Prepare(createCustomerQuery)
	if err != nil {
		return nil, err
	}
	createBranchQuery := `insert into BranchCompanies(
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
	createBranchStmt, err := db.Prepare(createBranchQuery)
	if err != nil {
		return nil, err
	}
	selectByNameQuery := `SELECT CodeNumber,
	NAME,
	DOI,
	GEMI,
	Phone,
	Mobile_Phone,
	Email,
	PostalCellName,
	PostalCellNumber,
	PostalCellPostalCode,
	PostalCellCity,
	AddStreet,
	AddNumber,
	AddPostalCode,
	AddCity,
	VatNumber,
	Country,
	Branch,
	Balance,
	Discount
	from customers  where NAME LIKE ? `
	selectByNameStmt, err := db.Prepare(selectByNameQuery)
	if err != nil {
		return nil, err
	}

	selectBranchQuery := `SELECT 
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
	selectBranchStmt, err := db.Prepare(selectBranchQuery)
	if err != nil {
		return nil, err
	}
	searchByIdQuery := `SELECT 
	CodeNumber,
	NAME,
	DOI,
	GEMI,
	Phone,
	Mobile_Phone,
	Email,
	PostalCellName,
	PostalCellNumber,
	PostalCellPostalCode,
	PostalCellCity,
	AddStreet,
	AddNumber,
	AddPostalCode,
	AddCity,
	VatNumber,
	Country,
	Branch,
	Balance,
	Discount 
	from customers  where CodeNumber== ? `
	searchByIdStmt, err := db.Prepare(searchByIdQuery)
	if err != nil {
		return nil, err
	}
	return &CustomerStmts{
		CreateCustomer:            insertStmt,
		CreateBranchCompany:       createBranchStmt,
		SearchByName:              selectByNameStmt,
		SearchBranch:              selectBranchStmt,
		SearchById:                searchByIdStmt,
		CustomerSuggestionsByName: customerSuggestionsStmt,
		BranchSuggestions:         branchSuggestionsStmt,
		SearchFullBranch:          searchFullBranchStmt,
	}, nil
}
