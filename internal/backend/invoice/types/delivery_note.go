package types

import (
	"-invoice_manager/internal/backend/invoice/payload"
)

type DeliveryNote struct {
	Payload *payload.InvoicePayload
}

func (r *DeliveryNote) Initialize() {
	r.Payload = &payload.InvoicePayload{}
	r.Payload.Invoices = make([]payload.Invoice, 1)
}
func (r *DeliveryNote) GetInvoice() (payload *payload.Invoice) {
	return &r.Payload.Invoices[0]
}

func (r *DeliveryNote) CalculateAlltheInvoiceLines() error {
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
