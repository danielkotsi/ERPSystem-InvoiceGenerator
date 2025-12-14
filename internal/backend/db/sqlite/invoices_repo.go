package sqlite

import (
	"bytes"
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/utils"
	"html/template"
	"log"
	"time"
)

type InvoiceRepo struct {
	DB *sql.DB
}

func NewInvoiceRepo(db *sql.DB) *InvoiceRepo {
	return &InvoiceRepo{DB: db}
}

func (r *InvoiceRepo) DesignInvoice(ctx context.Context, invo models.InvoicePayload) (pdf models.InvoicePayload, err error) {
	if err := r.CompleteInvoicePayload(ctx, &invo); err != nil {
		return models.InvoicePayload{}, err
	}

	// finalInvoice, err := r.AssembleFinalInvoice(ctx, &invo)
	// if err != nil {
	// 	return models.Invoice{}, err
	// }

	// pdf, err = r.MakePDF(ctx, &finalInvoice)
	// if err != nil {
	// 	return nil, err
	// }

	return invo, nil
}

func (r *InvoiceRepo) CompleteInvoicePayload(ctx context.Context, invo *models.InvoicePayload) error {
	if err := r.CompleteInvoiceHeader(&invo.Invoice.InvoiceHeader); err != nil {
		return err
	}
	// if err := r.CalculateAlltheInvoiceLines(invo.Invoice.InvoiceDetails, &invo.Invoice.InvoiceSummary); err != nil {
	// 	return err
	// }

	if err := r.CalculateIncomeClasiffication(&invo.Invoice.InvoiceSummary); err != nil {
		return err
	}

	return nil
}

func (r *InvoiceRepo) CalculateAlltheInvoiceLines(invoicelines []*models.InvoiceRow, summary *models.InvoiceSummary) error {
	for _, line := range invoicelines {
		if err := r.CalculateInvoiceLinePrices(line); err != nil {
			return err
		}
		summary.TotalNetValue += line.NetValue
		summary.TotalVatAmount += line.VatAmount
	}
	// summary.TotalWithVat = summary.TotalNetValue + summary.TotalVatAmount
	return nil
}

func (r *InvoiceRepo) CalculateInvoiceLinePrices(line *models.InvoiceRow) error {
	amount := map[int]float64{
		1: 0.24,
		2: 0.13,
	}
	line.NetValue = line.Quantity * line.UnitPrice
	line.VatAmount = line.Quantity * line.UnitPrice * amount[line.VatCategory]

	return nil
}

func (r *InvoiceRepo) CalculateIncomeClasiffication(*models.InvoiceSummary) error {
	//to do
	return nil
}

func (r *InvoiceRepo) CompleteInvoiceHeader(header *models.InvoiceHeader) error {
	if err := r.CalculateAA(header); err != nil {
		return err
	}
	header.IssueDate = time.Now().Format("2006-01-02")
	return nil
}
func (r *InvoiceRepo) CalculateAA(header *models.InvoiceHeader) error {
	// header.Aa = "12"
	return nil
}

func (r *InvoiceRepo) MakePostRequestToMyData(invo *models.InvoicePayload) error {
	//to do
	return nil
}

// we'll see if we are going to use this since i already have the makepostrequesttoMydata
func (r *InvoiceRepo) AssembleFinalInvoice(ctx context.Context, invo *models.InvoicePayload) (finalinvoice models.Invoice, err error) {

	return models.Invoice{}, err
}

func (r *InvoiceRepo) MakePDF(ctx context.Context, finalInvoice *models.InvoicePayload) (pdf []byte, err error) {
	tmpl, err := template.ParseFiles("../../assets/templates/invoice.page.html")
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]models.InvoicePayload{"Invoice": *finalInvoice})
	if err != nil {
		log.Println(err)
	}
	pdf, err = utils.HTMLtoPDF(buf.String())
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
