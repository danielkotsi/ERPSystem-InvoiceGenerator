package sqlite

import (
	"database/sql"
)

type InvoiceStmts struct {
	SaveSellingInvoice *sql.Stmt
	SaveBuyingInvoice  *sql.Stmt
	SaveDeliveryNote   *sql.Stmt
	SaveReciept        *sql.Stmt
	GetInvoiceInfo     *sql.Stmt
	UpdateBalance      *sql.Stmt
	UpdateAA           *sql.Stmt
	GetSellerInfo      *sql.Stmt
	GetBuyerBalance    *sql.Stmt
}

func NewInvoiceStmts(db *sql.DB) (*InvoiceStmts, error) {
	saveSelling_InvoiceQuery := ``
	saveSelling_InvoiceStmt, err := db.Prepare(saveSelling_InvoiceQuery)
	if err != nil {
		return nil, err
	}

	saveBuying_InvoiceQuery := ``
	saveBuying_InvoiceStmt, err := db.Prepare(saveBuying_InvoiceQuery)
	if err != nil {
		return nil, err
	}
	saveDeliveryNoteQuery := ``
	saveDeliveryNoteStmt, err := db.Prepare(saveDeliveryNoteQuery)
	if err != nil {
		return nil, err
	}
	saveRecieptQuery := ``
	saveRecieptStmt, err := db.Prepare(saveRecieptQuery)
	if err != nil {
		return nil, err
	}
	getInvoiceInfoquery := `select 
	users.CodeNumber, 
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
	series,
	aa
	from users 
	join user_invoice_types_series 
	on users.CodeNumber==user_invoice_types_series.codeNumber 
	where users.CodeNumber== ? 
	and invoice_type==?;
`
	getInvoiceInfoStmt, err := db.Prepare(getInvoiceInfoquery)
	if err != nil {
		return nil, err
	}
	updateBalanceQuery := `update customers 
	set Balance=? 
	where CodeNumber==?;`
	updateBalanceStmt, err := db.Prepare(updateBalanceQuery)
	if err != nil {
		return nil, err
	}
	updateAAquery := `update user_invoice_types_series set aa=? where invoice_type==?;`
	updateAAStmt, err := db.Prepare(updateAAquery)
	if err != nil {
		return nil, err
	}
	getSellerInfoQuery := `select 
	PostalCellName,
	PostalCellNumber,
	PostalCellPostalCode,
	PostalCellCity 
	from users where CodeNumber==?;`
	getSellerInfoStmt, err := db.Prepare(getSellerInfoQuery)
	if err != nil {
		return nil, err
	}
	getBuyerBalanceQuery := `select 
	Balance 
	from customers where CodeNumber=?;`
	getBuyerBalanceStmt, err := db.Prepare(getBuyerBalanceQuery)
	if err != nil {
		return nil, err
	}
	return &InvoiceStmts{
		SaveSellingInvoice: saveSelling_InvoiceStmt,
		SaveBuyingInvoice:  saveBuying_InvoiceStmt,
		SaveDeliveryNote:   saveDeliveryNoteStmt,
		SaveReciept:        saveRecieptStmt,
		GetInvoiceInfo:     getInvoiceInfoStmt,
		UpdateBalance:      updateBalanceStmt,
		UpdateAA:           updateAAStmt,
		GetSellerInfo:      getSellerInfoStmt,
		GetBuyerBalance:    getBuyerBalanceStmt,
	}, nil

}
