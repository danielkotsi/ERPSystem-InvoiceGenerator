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

type SellingInvoice struct {
	Payload *payload.InvoicePayload
	Logo    string
	Abspath string
}

func (r *SellingInvoice) Initialize() {
	r.Payload = &payload.InvoicePayload{}
	r.Payload.Invoices = make([]payload.Invoice, 1)
}

func (r *SellingInvoice) GetInvoice() (payload *payload.Invoice) {
	return &r.Payload.Invoices[0]
}

func (r *SellingInvoice) CalculateInvoiceLines() error {
	emptylines := 24
	invoicelines := r.GetInvoice().InvoiceDetails
	buyer := &r.GetInvoice().Byer
	summary := &r.GetInvoice().InvoiceSummary
	paymentmethods := r.GetInvoice().PaymentMethods
	for i, line := range invoicelines {
		emptylines--
		line.VatCategoryName = utils.VatNames(line.VatCategory)
		line.LineNumber = i + 1
		if err := r.InvoiceLinePrices(line, buyer.Discount); err != nil {
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
		if err := AddIncomeClassificationInSummary(line.IncomeClassification, summary); err != nil {
			return err
		}
	}
	summary.TotalGrossValue = utils.RoundTo2(summary.TotalNetValue + summary.TotalVatAmount)
	buyer.NewBalance = buyer.OldBalance

	if err := r.CompletePaymentMethods(paymentmethods, buyer, summary.TotalGrossValue); err != nil {
		return err
	}

	summary.Emptylines = make([]int, emptylines)
	return nil
}

func (r *SellingInvoice) InvoiceLinePrices(line *payload.InvoiceRow, discount int) error {
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

func (r *SellingInvoice) CompletePaymentMethods(paymentmethods *payload.PaymentMethods, buyer *payload.Company, totalgrossamount float64) error {
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

// makepdf must be in the domain interface
func (r *SellingInvoice) MakePDF(ctx context.Context) (pdf []byte, err error) {
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
