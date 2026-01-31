package sqlite

import (
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/payload"
	reposinterfaces "-invoice_manager/internal/backend/invoice/reposInterfaces"
	"-invoice_manager/internal/backend/invoice/types"
	"fmt"
	"strconv"
	"time"
)

type InvoiceRepo struct {
	DB    *sql.DB
	Stmts *InvoiceStmts
}

func NewInvoiceRepo(db *sql.DB, abspath string, logo string) (*InvoiceRepo, error) {
	InvoiceStmts, err := NewInvoiceStmts(db)
	if err != nil {
		return nil, err
	}
	return &InvoiceRepo{DB: db, Stmts: InvoiceStmts}, nil
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
	invoiceinfo.User.Address = &payload.AddressType{}

	err = r.Stmts.GetInvoiceInfo.
		QueryRowContext(ctx, "COMP01", string(invoicetype)).
		Scan(&invoiceinfo.User.CodeNumber,
			&invoiceinfo.User.Name,
			&invoiceinfo.User.DOI,
			&invoiceinfo.User.GEMI,
			&invoiceinfo.User.Phone,
			&invoiceinfo.User.Mobile_Phone,
			&invoiceinfo.User.Email,
			&invoiceinfo.User.PostalAddress.Naming,
			&invoiceinfo.User.PostalAddress.Cellnumber,
			&invoiceinfo.User.PostalAddress.PostalCode,
			&invoiceinfo.User.PostalAddress.City,
			&invoiceinfo.User.Address.Street,
			&invoiceinfo.User.Address.Number,
			&invoiceinfo.User.Address.PostalCode,
			&invoiceinfo.User.Address.City,
			&invoiceinfo.User.VatNumber,
			&invoiceinfo.User.Country,
			&invoiceinfo.User.Branch,
			&invoiceinfo.Invoiceinfo.Series,
			&invoiceinfo.Invoiceinfo.Aa)
	if err != nil {
		return invoiceinfo, err
	}

	err = r.CompleteHTMLinfo(&invoiceinfo, invoicetype)
	if err != nil {
		return invoiceinfo, err
	}
	return invoiceinfo, nil
}

func (r *InvoiceRepo) GetBuyerBalance(ctx context.Context, buyer *payload.Company) error {
	err := r.Stmts.GetBuyerBalance.
		QueryRowContext(ctx, buyer.CodeNumber).Scan(&buyer.OldBalance)
	if err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepo) GetSellerInfo(ctx context.Context, seller *payload.Company) error {
	//we might get an error here because of the comma inside scan
	err := r.Stmts.GetSellerInfo.
		QueryRowContext(ctx, seller.CodeNumber).
		Scan(
			&seller.PostalAddress.Naming,
			&seller.PostalAddress.Cellnumber,
			&seller.PostalAddress.PostalCode,
			&seller.PostalAddress.City,
		)
	if err != nil {
		return err
	}
	return nil
}
func (r *InvoiceRepo) UpdateBalance(ctx context.Context, buyerCodeNumber string, buyerNewBalance float64) error {
	if _, err := r.Stmts.UpdateBalance.
		ExecContext(ctx, buyerNewBalance, buyerCodeNumber); err != nil {
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

	if _, err := r.Stmts.UpdateAA.
		ExecContext(ctx, aa, invoicetype); err != nil {
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
	if err := r.UpdateDB(ctx, invoice.Byer.NewBalance, invoice.Byer.CodeNumber, invoice.InvoiceHeader.InvoiceType, invoice.InvoiceHeader.Aa); err != nil {
		return fmt.Errorf("Error on updating the Database %w", err)
	}
	// switch invo.(type) {
	// case *types.SellingInvoice:
	// 	_, err := r.Stmts.SaveSellingInvoice.ExecContext(ctx, invo.GetInvoice())
	// 	if err != nil {
	// 		return err
	// 	}
	// case *types.Buying_Invoice:
	// 	//this is so that it doesn't panic on empty values
	// 	invoice.Seller.Address = &payload.AddressType{}
	// 	invoice.Seller.Address.Street = "athina"
	// 	invoice.Seller.Address.Number = "123"
	// 	invoice.Seller.Address.PostalCode = "12fh3"
	// 	invoice.Seller.Address.City = "berling"
	// 	name := "alex"
	// 	invoice.Seller.Name = &name
	// 	invoice.Byer.Name = &name
	// 	invoice.InvoiceHeader.OtherDeliveryNoteHeader = &payload.OtherDeliveryNoteHeader{}
	// 	_, err := r.Stmts.SaveBuyingInvoice.ExecContext(ctx, invo.GetInvoice())
	// 	if err != nil {
	// 		return err
	// 	}
	// case *types.DeliveryNote:
	// 	_, err := r.Stmts.SaveDeliveryNote.ExecContext(ctx, invo.GetInvoice())
	// 	if err != nil {
	// 		return err
	// 	}
	// case *types.Reciept:
	// 	//this is so that it doesn't panic on empty values
	// 	invoice := invo.GetInvoice()
	// 	invoice.Seller.Address = &payload.AddressType{}
	// 	invoice.Seller.Address.Street = "athina"
	// 	invoice.Seller.Address.Number = "123"
	// 	invoice.Seller.Address.PostalCode = "12fh3"
	// 	invoice.Seller.Address.City = "berling"
	// 	name := "alex"
	// 	invoice.Seller.Name = &name
	// 	invoice.Byer.Name = &name
	// 	invoice.InvoiceHeader.OtherDeliveryNoteHeader = &payload.OtherDeliveryNoteHeader{}
	// 	_, err := r.Stmts.SaveReciept.ExecContext(ctx, invo.GetInvoice())
	// 	if err != nil {
	// 		return err
	// 	}
	// default:
	// }

	return nil
}
