package types

import (
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
)

type Reciept struct {
	Payload *payload.InvoicePayload
}

func (r *Reciept) GetInvoice() (payload *payload.Invoice) {
	return &r.Payload.Invoices[0]
}

func (r *Reciept) Initialize() {
	r.Payload = &payload.InvoicePayload{}
	r.Payload.Invoices = make([]payload.Invoice, 1)
}

func (r *Reciept) CalculateAlltheInvoiceLines() error {
	emptylines := 24
	invoicelines := r.GetInvoice().InvoiceDetails
	buyer := &r.GetInvoice().Byer
	summary := &r.GetInvoice().InvoiceSummary
	paymentmethods := r.GetInvoice().PaymentMethods
	for i, line := range invoicelines {
		emptylines--
		line.LineNumber = i + 1
		line.IncomeClassification.Amount = line.NetValue
		summary.TotalNetValue = line.NetValue
		if err := AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
			return err
		}
	}
	summary.TotalGrossValue = utils.RoundTo2(summary.TotalNetValue + summary.TotalVatAmount)
	buyer.NewBalance = buyer.OldBalance - summary.TotalGrossValue
	if err := r.CompletePaymentMethods(paymentmethods, buyer, summary.TotalGrossValue); err != nil {
		return err
	}
	summary.Emptylines = make([]int, emptylines)
	return nil
}

func (r *Reciept) CompletePaymentMethods(paymentmethods *payload.PaymentMethods, buyer *payload.Company, totalgrossamount float64) error {
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
		paymentmethods.Details[i].Amount = totalgrossamount
		if paymentmethods.Details[i].Type == 5 {
			buyer.NewBalance = utils.RoundTo2(buyer.OldBalance + totalgrossamount)
		}
	}

	return nil
}
