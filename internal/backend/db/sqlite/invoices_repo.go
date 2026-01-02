package sqlite

import (
	"bytes"
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/utils"
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"strconv"
	"time"
)

type InvoiceRepo struct {
	DB      *sql.DB
	abspath string
	logo    string
}

func NewInvoiceRepo(db *sql.DB, abspath string, logo string) *InvoiceRepo {
	return &InvoiceRepo{DB: db, abspath: abspath, logo: logo}
}

func (r *InvoiceRepo) AddToAA(ctx context.Context, invoicetype, aa string) error {
	aaint, err := strconv.Atoi(aa)
	if err != nil {
		return err
	}
	aaint++
	aa = fmt.Sprintf("%05d", aaint)
	fmt.Println(aa)
	query := `update user_invoice_types_series set aa=? where invoice_type==?;`

	if _, err := r.DB.ExecContext(ctx, query, aa, invoicetype); err != nil {
		return err
	}
	return nil
}
func (r *InvoiceRepo) CompleteInvoice(ctx context.Context, invo *models.Invoice) error {
	// invo.Seller.Address = nil
	if err := r.CompleteInvoiceHeader(&invo.InvoiceHeader); err != nil {
		return err
	}
	if err := r.CalculateAlltheInvoiceLines(invo.InvoiceHeader.InvoiceType, invo.InvoiceDetails, &invo.InvoiceSummary); err != nil {
		return err
	}

	return nil
}

func (r *InvoiceRepo) CalculateAlltheInvoiceLines(invoicetype string, invoicelines []*models.InvoiceRow, summary *models.InvoiceSummary) error {
	for i, line := range invoicelines {
		line.LineNumber = i + 1
		if invoicetype != "9.3" {
			if err := r.CalculateInvoiceLinePrices(line); err != nil {
				return err
			}
		}
		line.IncomeClassification.Amount = line.NetValue /* + line.VatAm unt */
		summary.TotalNetValue += line.NetValue
		summary.TotalNetValue = utils.RoundTo2(summary.TotalNetValue)
		summary.TotalVatAmount += line.VatAmount
		summary.TotalVatAmount = utils.RoundTo2(summary.TotalVatAmount)
		if err := r.AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
			return err
		}
	}
	summary.TotalGrossValue = summary.TotalNetValue + summary.TotalVatAmount
	summary.TotalGrossValue = utils.RoundTo2(summary.TotalGrossValue)
	return nil
}

func (r *InvoiceRepo) AddIncomeClassificationInSummary(classificationItem *models.ClassificationItem, summary *models.InvoiceSummary) error {
	index, exists := r.ClassificationCategoryExists(*classificationItem, summary.IncomeClassification)
	fmt.Println(classificationItem)
	if exists {
		summary.IncomeClassification[index].Amount += classificationItem.Amount
	} else {
		summary.IncomeClassification = append(summary.IncomeClassification, *classificationItem)
	}
	return nil
}

func (r *InvoiceRepo) ClassificationCategoryExists(classificationitem models.ClassificationItem, summary []models.ClassificationItem) (int, bool) {
	for index, category := range summary {
		if classificationitem.ClassificationCategory == category.ClassificationCategory && classificationitem.ClassificationType == category.ClassificationType {
			return index, true
		}
	}

	return 0, false
}

func (r *InvoiceRepo) CalculateInvoiceLinePrices(line *models.InvoiceRow) error {
	amount := map[int]float64{
		1:  0.24,
		2:  0.13,
		3:  0.06,
		4:  0.17,
		5:  0.09,
		6:  0.04,
		7:  0.00,
		8:  0.00,
		9:  0.03,
		10: 0.04,
	}
	line.NetValue = line.Quantity * line.UnitNetPrice
	line.VatAmount = line.Quantity * line.UnitNetPrice * amount[line.VatCategory]

	line.NetValue = utils.RoundTo2(line.NetValue)
	line.VatAmount = utils.RoundTo2(line.VatAmount)
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

	invoiceinfo.User.Address = &models.AddressType{}
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

func (r *InvoiceRepo) CompleteHTMLinfo(invoiceinfo *models.InvoiceHTMLinfo, invoicetype string) error {
	invoiceinfo.Invoiceinfo.Currency = "EUR"
	invoiceinfo.Invoiceinfo.Invoicetype = invoicetype
	switch invoicetype {
	case "1.1":
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_561_001"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category1_2"
		invoiceinfo.Invoiceinfo.MovePurpose = "1"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = true
	case "8.1":
		invoiceinfo.Invoiceinfo.IncomeClassificationType = ""
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = ""
		invoiceinfo.Invoiceinfo.IsDeliveryNote = false
	case "9.3":
		invoiceinfo.Invoiceinfo.IncomeClassificationType = ""
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category3"
		invoiceinfo.Invoiceinfo.MovePurpose = "3"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = true
	case "13.1":
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_201"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category2_2"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = false
	}
	return nil
}

func (r *InvoiceRepo) CompleteInvoiceHeader(header *models.InvoiceHeader) error {
	header.IssueDate = time.Now().Format("2006-01-02")
	header.Time = time.Now().Format("15:04")
	return nil
}

func (r *InvoiceRepo) MakePDF(ctx context.Context, finalInvoice *models.Invoice) (pdf []byte, err error) {
	// finalInvoice.QrBase64, err = utils.GenerateQRcodeBase64(finalInvoice.QrURL)
	finalInvoice.LogoImage = r.logo
	fmt.Println("this is the image base 64", finalInvoice.LogoImage)
	finalInvoice.QrBase64, err = utils.GenerateQRcodeBase64("http://localhost:8080")
	if err != nil {
		return nil, err
	}

	invoicehtmltemp := filepath.Join(r.abspath, "assets", "templates", "invoice.page.html")
	tmpl, err := template.ParseFiles(invoicehtmltemp)
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]models.Invoice{"Invoice": *finalInvoice})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(buf.String())
	pdf, err = utils.HTMLtoPDF2(buf.String())
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
