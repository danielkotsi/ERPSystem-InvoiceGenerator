package sqlite

import (
	"database/sql"
)

type CustomerStmts struct {
	CreateCustomer      *sql.Stmt
	CreateBranchCompany *sql.Stmt
	SearchByName        *sql.Stmt
	SearchBranch        *sql.Stmt
	SearchById          *sql.Stmt
}

func NewCustomerStmts(db *sql.DB) (*CustomerStmts, error) {
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
	return &CustomerStmts{CreateCustomer: insertStmt,
		CreateBranchCompany: createBranchStmt,
		SearchByName:        selectByNameStmt,
		SearchBranch:        selectBranchStmt,
		SearchById:          searchByIdStmt,
	}, nil
}
