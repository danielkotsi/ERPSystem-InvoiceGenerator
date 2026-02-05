package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"invoice_manager/internal/backend/invoice/models"
	"invoice_manager/internal/backend/invoice/payload"
	reposinterfaces "invoice_manager/internal/backend/invoice/reposInterfaces"
	"invoice_manager/internal/backend/invoice/types"
	"invoice_manager/internal/utils"
	"strconv"
	"time"
)

type InvoiceRepo struct {
	DB *sql.DB
}

func NewInvoiceRepo(db *sql.DB, abspath string, logo string) *InvoiceRepo {
	return &InvoiceRepo{DB: db}
}

func (r *InvoiceRepo) HydrateInvoice(ctx context.Context, invo reposinterfaces.Invoice_type) error {
	if err := r.GetSellerInfo(ctx, &invo.GetInvoice().Seller); err != nil {
		return err
	}
	if err := r.GetBuyerBalance(ctx, &invo.GetInvoice().Byer); err != nil {
		return err
	}
	if err := r.CompleteInvoiceHeader(&invo.GetInvoice().InvoiceHeader); err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepo) UpdateDB(ctx context.Context, buyerNewBalance float64, buyerCodeNumber, invoicetype, aa string) error {
	if err := r.UpdateBalance(ctx, buyerCodeNumber, buyerNewBalance); err != nil {
		return err
	}
	if err := r.AddToAA(ctx, invoicetype, aa); err != nil {
		return err
	}
	return nil
}
func (r *InvoiceRepo) GetInvoiceInfo(ctx context.Context, invoicetype types.InvoiceType) (invoiceinfo models.InvoiceHTMLinfo, err error) {
	query := `
select users.CodeNumber, 
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
from users join user_invoice_types_series on users.CodeNumber==user_invoice_types_series.codeNumber where users.CodeNumber== ? and invoice_type==?;
`
	rows, err := r.DB.QueryContext(ctx, query, "COMP01", string(invoicetype))
	if err != nil {
		return invoiceinfo, err
	}
	defer rows.Close()

	invoiceinfo.User.Address = &payload.AddressType{}
	for rows.Next() {
		if err := rows.Scan(&invoiceinfo.User.CodeNumber, &invoiceinfo.User.Name, &invoiceinfo.User.DOI, &invoiceinfo.User.GEMI, &invoiceinfo.User.Phone, &invoiceinfo.User.Mobile_Phone, &invoiceinfo.User.Email, &invoiceinfo.User.PostalAddress.Naming, &invoiceinfo.User.PostalAddress.Cellnumber, &invoiceinfo.User.PostalAddress.PostalCode, &invoiceinfo.User.PostalAddress.City, &invoiceinfo.User.Address.Street, &invoiceinfo.User.Address.Number, &invoiceinfo.User.Address.PostalCode, &invoiceinfo.User.Address.City, &invoiceinfo.User.VatNumber, &invoiceinfo.User.Country, &invoiceinfo.User.Branch, &invoiceinfo.Invoiceinfo.Series, &invoiceinfo.Invoiceinfo.Aa); err != nil {
			return invoiceinfo, err
		}
	}

	err = r.CompleteHTMLinfo(&invoiceinfo, invoicetype)
	if err != nil {
		return invoiceinfo, err
	}
	return invoiceinfo, nil
}

func (r *InvoiceRepo) GetBuyerBalance(ctx context.Context, buyer *payload.Company) error {
	query := "select Balance from customers where CodeNumber=?;"
	rows, err := r.DB.QueryContext(ctx, query, buyer.CodeNumber)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&buyer.OldBalance); err != nil {
			return err
		}
	}

	return nil
}
func (r *InvoiceRepo) GetSellerInfo(ctx context.Context, seller *payload.Company) error {
	query := `select PostalCellName, PostalCellNumber,PostalCellPostalCode, PostalCellCity from users where CodeNumber==?;`

	rows, err := r.DB.QueryContext(ctx, query, seller.CodeNumber)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&seller.PostalAddress.Naming, &seller.PostalAddress.Cellnumber, &seller.PostalAddress.PostalCode, &seller.PostalAddress.City); err != nil {
			return err
		}
	}
	return nil
}
func (r *InvoiceRepo) UpdateBalance(ctx context.Context, buyerCodeNumber string, buyerNewBalance float64) error {
	query := "update customers set Balance=? where CodeNumber==?;"
	if _, err := r.DB.ExecContext(ctx, query, utils.RoundTo2(buyerNewBalance), buyerCodeNumber); err != nil {
		return err
	}
	return nil
}
func (r *InvoiceRepo) AddToAA(ctx context.Context, invoicetype, aa string) error {
	aaint, err := strconv.Atoi(aa)
	if err != nil {
		return err
	}
	aaint++
	aa = fmt.Sprintf("%05d", aaint)
	query := `update user_invoice_types_series set aa=? where invoice_type==?;`

	if _, err := r.DB.ExecContext(ctx, query, aa, invoicetype); err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepo) CompleteHTMLinfo(invoiceinfo *models.InvoiceHTMLinfo, invoicetype types.InvoiceType) error {
	invoiceinfo.Invoiceinfo.Currency = "EUR"
	invoiceinfo.Invoiceinfo.Invoicetype = string(invoicetype)
	switch invoicetype {
	case types.SellingInvoiceType:
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_561_001"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category1_2"
		invoiceinfo.Invoiceinfo.MovePurpose = "1"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = true
	case types.RecieptInvoiceType:
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_561_001"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category1_8"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = false
		invoiceinfo.Invoiceinfo.VatCategory = 8
	case types.DeliveryNoteInvoiceType:
		invoiceinfo.Invoiceinfo.IncomeClassificationType = ""
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category3"
		invoiceinfo.Invoiceinfo.MovePurpose = "3"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = true
	case types.BuyingInvoiceType:
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_201"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category2_2"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = false
	}
	return nil
}

func (r *InvoiceRepo) CompleteInvoiceHeader(header *payload.InvoiceHeader) error {
	movepurpses := map[int]string{
		1: "Πώληση",
		3: "Δειγματισμός",
		9: "Αγορά",
	}
	header.IssueDate = time.Now().Format("2006-01-02")
	header.Time = time.Now().Format("15:04")
	header.MovePurposeName = movepurpses[header.MovePurpose]
	return nil
}

// over here i want to populate the address of the seller in the case of a reciept invoice
// becasue myData doesn't want it on the method SendInvoice but i need it on the final pdf
// and on the invoice that is going to be saved,
// the seller name,the Buyer name, the OtherDeliveryNoteHeader as well
func (r *InvoiceRepo) Save(ctx context.Context, invo reposinterfaces.Invoice_type) error {
	invoice := invo.GetInvoice()
	// uncomment the following lines for reciept invoice and buying invoice not to panic
	// invoice.Seller.Address = &payload.AddressType{}
	// invoice.Seller.Address.Street = "athina"
	// invoice.Seller.Address.Number = "123"
	// invoice.Seller.Address.PostalCode = "12fh3"
	// invoice.Seller.Address.City = "berling"
	// name := "alex"
	// invoice.Seller.Name = &name
	// invoice.Byer.Name = &name
	// invoice.InvoiceHeader.OtherDeliveryNoteHeader = &payload.OtherDeliveryNoteHeader{}
	if err := r.UpdateDB(ctx, invoice.Byer.NewBalance, invoice.Byer.CodeNumber, invoice.InvoiceHeader.InvoiceType, invoice.InvoiceHeader.Aa); err != nil {
		return fmt.Errorf("Error on updating the Database %w", err)
	}
	return nil
}
