package services

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/repos"
	"-invoice_manager/internal/utils"
	"encoding/xml"
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

func (s *InvoiceService) CreateInvoice(ctx context.Context, r *http.Request) (pdf []byte, err error) {
	var invo models.Invoice
	err = utils.ParseFormData(r, &invo)
	if err != nil {
		return nil, err
	}
	fmt.Println("hello this is the invo from the form", invo)

	err = s.Invoice.CompleteInvoice(ctx, &invo)
	if err != nil {
		return nil, err
	}

	xmlinvo, err := xml.MarshalIndent(invo, "", "  ")
	if err != nil {
		return nil, err
	}
	// err = s.MyData.SendInvoice(ctx, &invo)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("\n\n\nthis is the invo QrCodeURL \n\n", invo.QrURL)
	//
	// pdf, err = s.Invoice.MakePDF(ctx, &invo)
	// if err != nil {
	// 	return nil, err
	// }

	return xmlinvo, nil
}
