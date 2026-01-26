package invoice

import (
	"context"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/reposInterfaces"
	"-invoice_manager/internal/backend/invoice/types"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type InvoiceService struct {
	Invoice reposinterfaces.Invoice_repo
	MyData  MyData_repo
}

func NewInvoiceService(in reposinterfaces.Invoice_repo, mydata MyData_repo) *InvoiceService {
	return &InvoiceService{
		Invoice: in,
		MyData:  mydata,
	}
}

func (s *InvoiceService) GetInvoiceInfo(ctx context.Context, invoicetype types.InvoiceType) (invoiceinfo models.InvoiceHTMLinfo, err error) {
	invoiceinfo, err = s.Invoice.GetInvoiceInfo(ctx, invoicetype)
	if err != nil {
		return invoiceinfo, fmt.Errorf("Error In Getting InvoiceHTML info from DB: %w", err)
	}
	return invoiceinfo, nil
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, invo reposinterfaces.Invoice_type) (pdf []byte, err error) {
	if err := s.Invoice.HydrateInvoice(ctx, invo); err != nil {
		return nil, fmt.Errorf("Error in Hydration from DB: %w", err)
	}
	if err := invo.CalculateInvoiceLines(); err != nil {
		return nil, fmt.Errorf("Error in InvoiceLines Calculation: %w", err)
	}
	if err := s.MyData.SendInvoice(ctx, invo); err != nil {
		return nil, fmt.Errorf("Error Sending Invoice to Mydata: %w", err)
	}
	if err := s.Invoice.Save(ctx, invo); err != nil {
		return nil, fmt.Errorf("Error Saving the Invoice to the DB: %w", err)
	}
	data, err := json.Marshal(invo.GetInvoice())
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("product.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	pdf, err = invo.MakePDF(ctx)
	if err != nil {
		return nil, fmt.Errorf("Error in PDF Generation: %w", err)
	}
	return pdf, nil
}
