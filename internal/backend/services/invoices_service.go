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

func (s *InvoiceService) CreateInvoice(ctx context.Context, r *http.Request) (pdf []byte, err error) {
	var invo models.Invoice
	err = utils.ParseFormData(r, &invo)
	if err != nil {
		return nil, err
	}
	fmt.Println("hello this is the invo from the form", invo)
	invoicePayload, err := s.Invoice.DesignInvoice(ctx, invo)
	if err != nil {
		return pdf, err
	}

	completeinvo, err := s.MyData.SendInvoice(ctx, &invoicePayload)
	if err != nil {
		return completeinvo, err
	}
	// pdf, err := s.Invoice.MakePDF
	return completeinvo, nil
}
