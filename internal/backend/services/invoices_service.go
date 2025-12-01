package services

import (
	// "context"
	// "-invoice_manager/internal/backend/models"
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

func (s *InvoiceService) CreateInvoice(ctx context.Context, r *http.Request) (pdf models.InvoicePayload, err error) {
	var invo models.InvoicePayload
	err = utils.ParseFormData(r, &invo)
	if err != nil {
		return models.InvoicePayload{}, err
	}
	fmt.Println("this is the seller", invo.Invoice.Seller)
	fmt.Println("this is the seller", invo.Invoice.Seller)
	fmt.Println("this is the seller", invo.Invoice.Seller.Address.Street)
	fmt.Println("this is the buyer", invo.Invoice.Byer)
	fmt.Println("this is the buyer", *invo.Invoice.Byer.Address.Street)
	fmt.Println("this is the invoice", invo.Invoice.InvoiceDetails)

	invoicePayload, err := s.Invoice.DesignInvoice(ctx, invo)
	if err != nil {
		return pdf, err
	}

	if err := s.MyData.SendInvoice(ctx, &invoicePayload); err != nil {
	}
	return pdf, nil
}
