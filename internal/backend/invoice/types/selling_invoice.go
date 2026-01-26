package types

import (
	"bytes"
	"context"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
	"fmt"
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
	// r.GetInvoice().QrBase64, err = utils.GenerateQRcodeBase64(r.GetInvoice().QrURL)
	// r.GetInvoice().LogoImage = r.Logo
	// if err != nil {
	// 	return nil, err
	// }

	pdf, err := r.GeneratePDFfromTemp()
	if err != nil {
		return nil, err
	}

	qrpng, err := qrcode.Encode("http://localhost:8080", qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	pdf.ImageFromImageInBytes(qrpng, 20, 20, &gopdf.Rect{
		W: 80,
		H: 80,
	})
	err = pdf.AddTTFFont("OpenSans", "../../../../../usr/share/fonts/open-sans/OpenSans-Bold.ttf")
	if err != nil {
		log.Fatal(err)
	}
	err = pdf.SetFont("OpenSans", "", 14)
	if err != nil {
		log.Fatal(err)
	}
	pdf.SetXY(500, 50)

	pdf.MultiCell(&gopdf.Rect{
		W: 300, // wrap at 300 points
		H: 16,  // 16pt line height
	}, "This is a very long product description that will wrap automatically inside the defined rectangle.")
	result := &bytes.Buffer{}
	_, err = pdf.WriteTo(result)
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

func (r *SellingInvoice) GeneratePDFfromTemp() (*gopdf.GoPdf, error) {
	pdf := &gopdf.GoPdf{}

	// Start the PDF (normal start)
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	// Import the first page of the template
	err := pdf.ImportPagesFromSource(r.PDFtemplate, "/MediaBox")
	if err != nil {
		return nil, fmt.Errorf("couldn't load template into pdf %w", err)
	}

	return pdf, nil
}
