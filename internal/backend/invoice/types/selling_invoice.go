package types

import (
	"bytes"
	"context"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
	"log"

	"github.com/signintech/gopdf"
	"github.com/skip2/go-qrcode"
)

type SellingInvoice struct {
	Payload     *payload.InvoicePayload
	Logo        string
	Abspath     string
	PDFtemplate []byte
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
func (r *SellingInvoice) MakePDF(ctx context.Context) (resultpdf []byte, err error) {
	invo := r.GetInvoice()
	invo.QrBase64, err = utils.GenerateQRcodeBase64(r.GetInvoice().QrURL)
	invo.LogoImage = r.Logo
	if err != nil {
		return nil, err
	}

	pdf, err := GeneratePDFfromTemp()
	if err != nil {
		return nil, err
	}

	qrpng, err := qrcode.Encode(invo.QrURL, qrcode.Medium, 256)
	if err != nil {
		log.Fatal(err)
	}
	pdf.ImageFromImageFile(invo.LogoImage, 35, 15, &gopdf.Rect{
		W: 155,
		H: 95,
	})
	pdf.ImageFromImageInBytes(qrpng, 480, 15, &gopdf.Rect{
		W: 100,
		H: 100,
	})
	err = pdf.AddTTFFont("OpenSans", "/usr/share/fonts/open-sans/OpenSans-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	err = pdf.AddTTFFont("OpenSansBold", "../../../../../usr/share/fonts/open-sans/OpenSans-Bold.ttf")
	if err != nil {
		log.Fatal(err)
	}
	MakeHeader(pdf, invo)
	MakeVatCalculations(pdf, invo.InvoiceSummary)
	MakePrices(pdf, invo.InvoiceSummary)
	MakeInvoiceHeader(pdf, invo)
	MakeBalance(pdf, invo)
	MakeByer(pdf, invo.Byer)
	MakeDelivery(pdf, invo)
	MakeDetails(pdf, invo.InvoiceDetails)

	result := &bytes.Buffer{}
	_, err = pdf.WriteTo(result)
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}
