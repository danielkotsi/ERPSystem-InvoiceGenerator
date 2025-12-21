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
	"time"
)

type InvoiceRepo struct {
	DB      *sql.DB
	abspath string
}

func NewInvoiceRepo(db *sql.DB, abspath string) *InvoiceRepo {
	return &InvoiceRepo{DB: db, abspath: abspath}
}

func (r *InvoiceRepo) CompleteInvoice(ctx context.Context, invo *models.Invoice) error {
	// invo.Seller.Address = nil
	if err := r.CompleteInvoiceHeader(&invo.InvoiceHeader); err != nil {
		return err
	}
	if err := r.CalculateAlltheInvoiceLines(invo.InvoiceDetails, &invo.InvoiceSummary); err != nil {
		return err
	}

	return nil
}

func (r *InvoiceRepo) CalculateAlltheInvoiceLines(invoicelines []*models.InvoiceRow, summary *models.InvoiceSummary) error {
	for i, line := range invoicelines {
		line.LineNumber = i + 1
		if err := r.CalculateInvoiceLinePrices(line); err != nil {
			return err
		}
		summary.TotalNetValue += line.NetValue
		summary.TotalVatAmount += line.VatAmount
		if err := r.AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
			return err
		}
	}
	summary.TotalGrossValue = summary.TotalNetValue + summary.TotalVatAmount
	return nil
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

	return nil
}

func (r *InvoiceRepo) AddIncomeClassificationInSummary(classificationItem models.ClassificationItem, summary *models.InvoiceSummary) error {
	index, exists := r.ClassificationCategoryExists(classificationItem, summary.IncomeClassification)
	if exists {
		summary.IncomeClassification[index].Amount += classificationItem.Amount
	} else {
		summary.IncomeClassification = append(summary.IncomeClassification, classificationItem)
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

func (r *InvoiceRepo) CompleteInvoiceHeader(header *models.InvoiceHeader) error {
	if err := r.CalculateAA(header); err != nil {
		return err
	}
	header.IssueDate = time.Now().Format("2006-01-02")
	return nil
}
func (r *InvoiceRepo) CalculateAA(header *models.InvoiceHeader) error {
	header.Aa = "12"
	return nil
}

func (r *InvoiceRepo) MakePDF(ctx context.Context, finalInvoice *models.Invoice) (pdf []byte, err error) {
	finalInvoice.QrBase64, err = utils.GenerateQRcodeBase64(finalInvoice.QrURL)
	if err != nil {
		return nil, err
	}

	invoicehtmltemp := filepath.Join(r.abspath, "assets", "templates", "invoice.page.html")
	tmpl, err := template.ParseFiles(invoicehtmltemp)
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	fmt.Println("this is the finalinvoice MARK", finalInvoice.MARK)
	err = tmpl.Execute(&buf, map[string]models.Invoice{"Invoice": *finalInvoice})
	if err != nil {
		log.Println(err)
	}

	pdf, err = utils.HTMLtoPDF(buf.String())
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
