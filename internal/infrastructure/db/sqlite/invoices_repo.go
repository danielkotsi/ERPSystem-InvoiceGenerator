package sqlite

import (
	"bytes"
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
	"html/template"
	"log"
	"path/filepath"
)

type InvoiceRepo struct {
	DB      *sql.DB
	abspath string
	logo    string
}

func NewInvoiceRepo(db *sql.DB, abspath string, logo string) *InvoiceRepo {
	return &InvoiceRepo{DB: db, abspath: abspath, logo: logo}
}

func (r *InvoiceRepo) CompleteInvoice(ctx context.Context, invo *payload.Invoice) error {
	if err := r.GetSellerInfo(ctx, &invo.Seller); err != nil {
		return err
	}
	if err := r.GetBuyerBalance(ctx, &invo.Byer); err != nil {
		return err
	}
	if err := r.CompleteInvoiceHeader(&invo.InvoiceHeader); err != nil {
		return err
	}
	if err := r.CalculateAlltheInvoiceLines(invo.InvoiceHeader.InvoiceType, invo.PaymentMethods, invo.InvoiceDetails, &invo.InvoiceSummary, &invo.Byer); err != nil {
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
func (r *InvoiceRepo) GetInvoiceInfo(ctx context.Context, invoicetype string) (invoiceinfo models.InvoiceHTMLinfo, err error) {
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
	rows, err := r.DB.QueryContext(ctx, query, "COMP01", invoicetype)
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

func (r *InvoiceRepo) MakePDF(ctx context.Context, finalInvoice *payload.Invoice) (pdf []byte, err error) {
	finalInvoice.QrBase64, err = utils.GenerateQRcodeBase64(finalInvoice.QrURL)
	finalInvoice.LogoImage = r.logo
	if err != nil {
		return nil, err
	}

	var invoicehtmltemp string
	switch finalInvoice.InvoiceHeader.InvoiceType {
	case "8.1":
		invoicehtmltemp = filepath.Join(r.abspath, "assets", "templates", "reciept_invoice.page.html")
	default:
		invoicehtmltemp = filepath.Join(r.abspath, "assets", "templates", "invoice.page.html")
	}
	tmpl, err := template.ParseFiles(invoicehtmltemp)
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]payload.Invoice{"Invoice": *finalInvoice})
	if err != nil {
		log.Println(err)
	}

	pdf, err = utils.HTMLtoPDF2(buf.String())
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
