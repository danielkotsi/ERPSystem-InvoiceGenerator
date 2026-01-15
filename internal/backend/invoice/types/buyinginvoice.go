package types

import (
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
)

type Buying_Invoice struct {
	Payload *payload.InvoicePayload
}

func (r *Buying_Invoice) CalculateAlltheInvoiceLines(invoicetype string, paymentmethods *payload.PaymentMethods, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company) error {
	if err := r.BuyingInvoiceLines(invoicetype, invoicelines, summary, buyer, paymentmethods); err != nil {
		return err
	}
	return nil
}

func (r *Buying_Invoice) BuyingInvoiceLines(invoicetype string, invoicelines []*payload.InvoiceRow, summary *payload.InvoiceSummary, buyer *payload.Company, paymentmethods *payload.PaymentMethods) error {
	emptylines := 24
	for i, line := range invoicelines {
		emptylines--
		line.VatCategoryName = utils.VatNames(line.VatCategory)
		line.LineNumber = i + 1
		if err := r.InvoiceLinePrices(line, buyer.Discount); err != nil {
			return err
		}
		line.ExpensesClassification.Amount = line.NetValue /* + line.VatAm unt */
		summary.TotalDiscount += line.DiscountAmount
		summary.TotalNetBeforeDiscount += line.TotalNetBeforeDiscount
		summary.TotalNetBeforeDiscount = utils.RoundTo2(summary.TotalNetBeforeDiscount)
		summary.TotalNetValue += line.NetValue
		summary.TotalNetValue = utils.RoundTo2(summary.TotalNetValue)
		summary.TotalVatAmount += line.VatAmount
		summary.TotalVatAmount = utils.RoundTo2(summary.TotalVatAmount)
		if err := r.AddExpenseClassificationInSummary(line.ExpensesClassification, summary); err != nil {
			return err
		}
	}
	summary.TotalGrossValue = utils.RoundTo2(summary.TotalNetValue + summary.TotalVatAmount)
	summary.Emptylines = make([]int, emptylines)
	return nil
}

func (r *Buying_Invoice) InvoiceLinePrices(line *payload.InvoiceRow, discount int) error {
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

func (r *Buying_Invoice) AddExpenseClassificationInSummary(classificationItem *payload.ExpensesClassificationItem, summary *payload.InvoiceSummary) error {
	index, exists := ExpenseCategoryExists(*classificationItem, summary.ExpensesClassification)
	if exists {
		summary.ExpensesClassification[index].Amount += classificationItem.Amount
	} else {
		summary.ExpensesClassification = append(summary.ExpensesClassification, *classificationItem)
	}
	return nil
}

func (r *Buying_Invoice) CompletePaymentMethods(paymentmethods *payload.PaymentMethods, buyer *payload.Company, totalgrossamount float64) error {
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
