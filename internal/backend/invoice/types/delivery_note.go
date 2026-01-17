package types

import (
	"bytes"
	"context"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
	"html/template"
	"log"
	"path/filepath"
)

type DeliveryNote struct {
	Payload *payload.InvoicePayload
	Logo    string
	Abspath string
}

func (r *DeliveryNote) Initialize() {
	r.Payload = &payload.InvoicePayload{}
	r.Payload.Invoices = make([]payload.Invoice, 1)
}
func (r *DeliveryNote) GetInvoice() (payload *payload.Invoice) {
	return &r.Payload.Invoices[0]
}

func (r *DeliveryNote) CalculateInvoiceLines() error {
	emptylines := 24
	invoicelines := r.GetInvoice().InvoiceDetails
	buyer := &r.GetInvoice().Byer
	summary := &r.GetInvoice().InvoiceSummary
	for i, line := range invoicelines {
		emptylines--
		line.LineNumber = i + 1
		if err := AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
			return err
		}
	}
	buyer.NewBalance = buyer.OldBalance
	summary.Emptylines = make([]int, emptylines)
	return nil
}

func (r *DeliveryNote) MakePDF(ctx context.Context) (pdf []byte, err error) {
	r.GetInvoice().QrBase64, err = utils.GenerateQRcodeBase64(r.GetInvoice().QrURL)
	r.GetInvoice().LogoImage = r.Logo
	if err != nil {
		return nil, err
	}

	invoicehtmltemp := filepath.Join(r.Abspath, "assets", "templates", "invoice.page.html")
	tmpl, err := template.ParseFiles(invoicehtmltemp)
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]payload.Invoice{"Invoice": *r.GetInvoice()})
	if err != nil {
		log.Println(err)
	}

	pdf, err = utils.HTMLtoPDF2(buf.String())
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
