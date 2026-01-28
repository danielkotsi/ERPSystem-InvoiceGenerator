package types

import (
	"bytes"
	"context"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/utils"
	"github.com/signintech/gopdf"
	"github.com/skip2/go-qrcode"
	"log"
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

func (r *DeliveryNote) MakePDF(ctx context.Context) (resultpdf []byte, err error) {
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
