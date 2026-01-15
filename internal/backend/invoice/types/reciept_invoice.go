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

func (r *Reciept) CalculateAlltheInvoiceLines(invoicetype string, paymentmethods *payload.PaymentMethods, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company) error {
	if err := r.RecieptInvoiceLines(invoicetype, invoicelines, summary, buyer, paymentmethods); err != nil {
		return err
	}
	return nil
}

func (r *Reciept) RecieptInvoiceLines(invoicetype string, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company, paymentmethods *payload.PaymentMethods) error {
	emptylines := 24
	for i, line := range invoicelines {
		emptylines--
		line.LineNumber = i + 1
		line.IncomeClassification.Amount = line.NetValue /* + line.VatAm unt */
		summary.TotalNetValue = line.NetValue
		if err := r.AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
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

func (r *Reciept) InvoiceLinePrices(line *payload.InvoiceRow, discount int) error {
	line.Discount = float64(discount)
	floatdiscount := float64(discount) / 100

	totalNetPriceBeforeDiscount := *line.Quantity * line.UnitNetPrice
	line.DiscountAmount = utils.RoundTo2(totalNetPriceBeforeDiscount * floatdiscount)
	totalNetPriceAfterDiscount := totalNetPriceBeforeDiscount - line.DiscountAmount
	vatAfterDiscount := totalNetPriceAfterDiscount * utils.Vat(line.VatCategory)

	line.TotalNetBeforeDiscount = utils.RoundTo2(totalNetPriceBeforeDiscount)
	line.NetValue = utils.RoundTo2(totalNetPriceAfterDiscount)
	line.VatAmount = utils.RoundTo2(vatAfterDiscount)

	return nil
}

func (r *Reciept) AddIncomeClassificationInSummary(classificationItem *payload.ClassificationItem, summary *payload.InvoiceSummary) error {
	index, exists := IncomeCategoryExists(*classificationItem, summary.IncomeClassification)
	if exists {
		summary.IncomeClassification[index].Amount += classificationItem.Amount
	} else {
		summary.IncomeClassification = append(summary.IncomeClassification, *classificationItem)
	}
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
