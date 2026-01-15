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

func (r *DeliveryNote) CalculateAlltheInvoiceLines(invoicetype string, paymentmethods *payload.PaymentMethods, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company) error {
	emptylines := 24
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
