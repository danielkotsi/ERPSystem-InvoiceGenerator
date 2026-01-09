package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
	"-invoice_manager/internal/utils"
	"fmt"
	"net/http"
)

type InvoiceService struct {
	Invoice repository.Invoice_repo
	MyData  repository.MyData_repo
}

func NewInvoiceService(in repository.Invoice_repo, mydata repository.MyData_repo) *InvoiceService {
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
	var invo models.Invoice
	err = utils.ParseFormData(r, &invo)
	if err != nil {
		return nil, err
	}

	err = s.Invoice.CompleteInvoice(ctx, &invo)
	if err != nil {
		return nil, err
	}

	fmt.Println(invo.PaymentMethods.Details[0].Type)
	fmt.Println(invo.PaymentMethods.Details[0].Name)
	// pdf, err = xml.MarshalIndent(invo, "", "  ")
	// if err != nil {
	// 	return nil, err
	// }
	// err = s.MyData.SendInvoice(ctx, &invo)
	// if err != nil {
	// 	return nil, err
	// }
	//
	// err = s.Invoice.AddToAA(ctx, invo.InvoiceHeader.InvoiceType, invo.InvoiceHeader.Aa)
	// if err != nil {
	// 	return nil, err
	// }
	pdf, err = s.Invoice.MakePDF(ctx, &invo)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
