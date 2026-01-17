package adapter

import (
	"-invoice_manager/internal/backend/invoice/reposInterfaces"
	"-invoice_manager/internal/backend/invoice/types"
	"-invoice_manager/internal/utils"
	"errors"
	"net/http"
)

// logo might go away from this and be fetched from the db for each user
type InvoiceParser struct {
	Abspath string
	Logo    string
}

func NewInvoiceParser(logo string, abspath string) *InvoiceParser {
	return &InvoiceParser{Logo: logo, Abspath: abspath}
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
