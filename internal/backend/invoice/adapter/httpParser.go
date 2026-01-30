package adapter

import (
	"-invoice_manager/internal/backend/invoice/reposInterfaces"
	"-invoice_manager/internal/backend/invoice/types"
	"-invoice_manager/internal/utils"
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// logo might go away from this and be fetched from the db for each user
type InvoiceParser struct {
	Abspath     string
	Logo        string
	PDFTemplate []byte
	//this is where the template for the pdf is gonna live and will be included inside the invoice structs
	//so that the invoices can use it in the makePDF function
	//the fonts must be there as well so that nothing needs to be loaded at runtime, everything in start up
}

func NewInvoiceParser(logo string, abspath string) *InvoiceParser {

	pdftemplate := filepath.Join(abspath, "assets", "pdftemplates", "invoicetemplate.pdf")
	data, err := os.ReadFile(pdftemplate)
	if err != nil {
		log.Fatalf("Failed to read template: %v", err)
	}
	return &InvoiceParser{Logo: logo, Abspath: abspath, PDFTemplate: data}
}
func (a *InvoiceParser) GetInvoiceTypeFromParameter(r *http.Request) types.InvoiceType {
	return types.InvoiceType(r.URL.Query().Get("invoice_type"))
}

func (a *InvoiceParser) ParseInvoiceFromRequest(r *http.Request) (invoicetype reposinterfaces.Invoice_type, err error) {
	switch types.InvoiceType(r.FormValue("invoiceHeader.invoiceType")) {
	case types.SellingInvoiceType:
		invoice := &types.SellingInvoice{}
		invoice.Initialize()
		err = utils.ParseFormData(r, &invoice.Payload.Invoices[0])
		if err != nil {
			return nil, err
		}
		invoice.Abspath = a.Abspath
		invoice.Logo = a.Logo
		invoice.PDFtemplate = a.PDFTemplate
		return invoice, nil
	case types.BuyingInvoiceType:
		invoice := &types.Buying_Invoice{}
		invoice.Initialize()
		err = utils.ParseFormData(r, &invoice.Payload.Invoices[0])
		if err != nil {
			return nil, err
		}
		invoice.Abspath = a.Abspath
		invoice.Logo = a.Logo
		return invoice, nil
	case types.DeliveryNoteInvoiceType:
		invoice := &types.DeliveryNote{}
		invoice.Initialize()
		err = utils.ParseFormData(r, &invoice.Payload.Invoices[0])
		if err != nil {
			return nil, err
		}
		invoice.Abspath = a.Abspath
		invoice.Logo = a.Logo
		return invoice, nil
	case types.RecieptInvoiceType:
		invoice := &types.Reciept{}
		invoice.Initialize()
		err = utils.ParseFormData(r, &invoice.Payload.Invoices[0])
		if err != nil {
			return nil, err
		}
		invoice.Abspath = a.Abspath
		invoice.Logo = a.Logo
		return invoice, nil
	default:
		return nil, errors.New("InvoiceType Not Supperted or Invalid InvoiceType")

	}
}
