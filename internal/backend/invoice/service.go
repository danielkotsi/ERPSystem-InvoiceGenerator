package invoice

import (
	"context"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/backend/invoice/types"
	"-invoice_manager/internal/utils"
	"errors"
	"net/http"
)

type InvoiceService struct {
	Invoice Invoice_repo
	MyData  MyData_repo
}

func NewInvoiceService(in Invoice_repo, mydata MyData_repo) *InvoiceService {
	return &InvoiceService{
		Invoice: in,
		MyData:  mydata,
	}
}

func (s *InvoiceService) GetInvoiceInfo(ctx context.Context, r *http.Request) (invoiceinfo models.InvoiceHTMLinfo, invoiceHTML string, err error) {
	invoicetype := r.URL.Query().Get("invoice_type")

	invoiceTypes := map[string]string{
		"1.1":  "create_selling_invoice.page.html",
		"13.1": "create_buying_invoice.page.html",
		"9.3":  "create_deliverynote_invoice.page.html",
		"8.1":  "create_reciept_invoice.page.html",
	}
	invoiceinfo, err = s.Invoice.GetInvoiceInfo(ctx, invoicetype)
	if err != nil {
		return invoiceinfo, invoiceTypes[invoicetype], err
	}
	return invoiceinfo, invoiceTypes[invoicetype], nil
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, r *http.Request) (pdf []byte, err error) {
	invo, err := s.ParseFormIntoInvoiceType(r)
	if err != nil {
		return nil, err
	}
	err = s.Invoice.CompleteInvoice(ctx, invo.GetInvoice())
	if err != nil {
		return nil, err
	}
	err = invo.CalculateAlltheInvoiceLines()
	if err != nil {
		return nil, err
	}
	err = s.MyData.SendInvoice(ctx, invo.GetInvoice())
	if err != nil {
		return nil, err
	}
	err = s.Invoice.UpdateDB(ctx, invo.GetInvoice().Byer.NewBalance, invo.GetInvoice().Byer.CodeNumber, invo.GetInvoice().InvoiceHeader.InvoiceType, invo.GetInvoice().InvoiceHeader.Aa)
	if err != nil {
		return nil, err
	}
	pdf, err = s.Invoice.MakePDF(ctx, invo.GetInvoice())
	if err != nil {
		return nil, err
	}

	return pdf, nil
}

func (s *InvoiceService) ParseFormIntoInvoiceType(r *http.Request) (invoice Invoice_type, err error) {
	switch r.FormValue("invoiceHeader.invoiceType") {
	case "1.1":
		invoice := &types.SellingInvoice{}
		invoice.Payload = &payload.InvoicePayload{}
		invoice.Payload.Invoices = make([]payload.Invoice, 1)
		err = utils.ParseFormData(r, &invoice.Payload.Invoices[0])
		if err != nil {
			return nil, err
		}
		return invoice, nil
	case "13.1":
	case "9.3":
	case "8.1":
	default:

	}
	return nil, errors.New("Invalid Invoice Type")
}
