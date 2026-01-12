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

func (r *InvoiceRepo) CompleteInvoice(ctx context.Context, invo *models.Invoice) error {
	// invo.Seller.Address = nil
	if err := r.GetSellerInfo(ctx, &invo.Seller); err != nil {
		return err
	}
	if err := r.GetBuyerBalance(ctx, &invo.Byer); err != nil {
		return err
	}
	if err := r.CompleteInvoiceHeader(&invo.InvoiceHeader); err != nil {
		return err
	}
	if err := r.CalculateAlltheInvoiceLines(invo.InvoiceHeader.InvoiceType, invo.InvoiceDetails, &invo.InvoiceSummary, &invo.Byer); err != nil {
		return err
	}
	if invo.PaymentMethods != nil {
		if err := r.CompletePaymentMethods(ctx, invo.PaymentMethods); err != nil {
			return err
		}
	}

	return nil
}

func (r *InvoiceRepo) GetBuyerBalance(ctx context.Context, buyer *models.Company) error {
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
func (r *InvoiceRepo) CompletePaymentMethods(ctx context.Context, paymentmethods *models.PaymentMethods) error {
	paymenttypes := map[string]int{
		"Επαγ. Λογαριασμός Πληρωμών Ημεδαπής":  1,
		"Επαγ. Λογαριασμός Πληρωμών Αλλοδαπής": 2,
		"Μετρητά":              3,
		"Επιταγή":              4,
		"Επί Πιστώσει":         5,
		"Web Banking":          6,
		"POS / e-POS":          7,
		"Άμεσες Πληρωμές IRIS": 8,
	}
	for i, payment := range paymentmethods.Details {
		paymentmethods.Details[i].Type = paymenttypes[payment.Name]
	}

	return nil
}
func (r *InvoiceRepo) GetSellerInfo(ctx context.Context, seller *models.Company) error {
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

	fmt.Println("hello this is the code", seller.CodeNumber)
	fmt.Println("hello this is the postall cell name", seller.PostalAddress.Naming)
	return nil
}
func (r *InvoiceRepo) CalculateAlltheInvoiceLines(invoicetype string, invoicelines []*models.InvoiceRow, summary *models.InvoiceSummary, buyer *models.Company) error {
	vatNames := map[int]int{
		1:  24,
		2:  13,
		3:  6,
		4:  17,
		5:  9,
		6:  4,
		7:  0,
		8:  0,
		9:  3,
		10: 4,
	}
	emptylines := 24
	for i, line := range invoicelines {
		emptylines--
		line.VatCategoryName = vatNames[line.VatCategory]
		line.LineNumber = i + 1
		if invoicetype != "9.3" && invoicetype != "8.1" {
			if err := r.CalculateInvoiceLinePrices(line, buyer.Discount); err != nil {
				return err
			}
			line.IncomeClassification.Amount = line.NetValue /* + line.VatAm unt */

			summary.TotalDiscount += line.DiscountAmount

			summary.TotalNetBeforeDiscount += line.TotalNetBeforeDiscount
			summary.TotalNetBeforeDiscount = utils.RoundTo2(summary.TotalNetBeforeDiscount)

			summary.TotalNetValue += line.NetValue
			summary.TotalNetValue = utils.RoundTo2(summary.TotalNetValue)
			summary.TotalVatAmount += line.VatAmount
			summary.TotalVatAmount = utils.RoundTo2(summary.TotalVatAmount)
		}
		if invoicetype == "8.1" {
			line.NetValue = 2.0
			line.IncomeClassification.Amount = line.NetValue /* + line.VatAm unt */
			summary.TotalNetValue = 2.0
		}
		if err := r.AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
			return err
		}
	}
	summary.TotalGrossValue = utils.RoundTo2(summary.TotalNetValue + summary.TotalVatAmount)
	buyer.NewBalance = buyer.OldBalance + summary.TotalGrossValue
	summary.Emptylines = make([]int, emptylines)
	return nil
}

func (r *InvoiceRepo) AddIncomeClassificationInSummary(classificationItem *models.ClassificationItem, summary *models.InvoiceSummary) error {
	index, exists := r.ClassificationCategoryExists(*classificationItem, summary.IncomeClassification)
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

func (r *InvoiceRepo) CalculateInvoiceLinePrices(line *models.InvoiceRow, discount int) error {
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
	line.Discount = float64(discount)
	floatdiscount := float64(discount) / 100

	totalNetPriceBeforeDiscount := *line.Quantity * line.UnitNetPrice
	line.DiscountAmount = utils.RoundTo2(totalNetPriceBeforeDiscount * floatdiscount)
	totalNetPriceAfterDiscount := totalNetPriceBeforeDiscount - line.DiscountAmount
	// vatBeforeDiscount := totalNetPriceBeforeDiscount * amount[line.VatCategory]
	vatAfterDiscount := totalNetPriceAfterDiscount * amount[line.VatCategory]
	// line.TotalAfterDiscount = totalNetPriceAfterDiscount

	line.TotalNetBeforeDiscount = utils.RoundTo2(totalNetPriceBeforeDiscount)
	line.NetValue = utils.RoundTo2(totalNetPriceAfterDiscount)
	line.VatAmount = utils.RoundTo2(vatAfterDiscount)

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
func (r *InvoiceRepo) UpdateBalance(ctx context.Context, buyerCodeNumber string, buyerNewBalance float64) error {
	query := "update customers set Balance=? where CodeNumber==?;"
	if _, err := r.DB.ExecContext(ctx, query, buyerNewBalance, buyerCodeNumber); err != nil {
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
	fmt.Println(aa)
	query := `update user_invoice_types_series set aa=? where invoice_type==?;`

	if _, err := r.DB.ExecContext(ctx, query, aa, invoicetype); err != nil {
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
		invoiceinfo.Invoiceinfo.IncomeClassificationType = "E3_561_001"
		invoiceinfo.Invoiceinfo.IncomeClassificationCat = "category1_8"
		invoiceinfo.Invoiceinfo.IsDeliveryNote = false
		invoiceinfo.Invoiceinfo.VatCategory = 8
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

func (r *InvoiceRepo) MakePDF(ctx context.Context, finalInvoice *models.Invoice) (pdf []byte, err error) {
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
	err = tmpl.Execute(&buf, map[string]models.Invoice{"Invoice": *finalInvoice})
	if err != nil {
		log.Println(err)
	}

	pdf, err = utils.HTMLtoPDF2(buf.String())
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
