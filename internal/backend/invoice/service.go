package invoice

import (
	"context"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/reposInterfaces"
	"-invoice_manager/internal/backend/invoice/types"
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
		return invoiceinfo, err
	}
	return invoiceinfo, nil
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, invo reposinterfaces.Invoice_type) (pdf []byte, err error) {
	err = s.Invoice.CompleteInvoice(ctx, invo)
	if err != nil {
		return nil, err
	}
	err = invo.CalculateAlltheInvoiceLines()
	if err != nil {
		return nil, err
	}
	err = s.MyData.SendInvoice(ctx, invo)
	if err != nil {
		return nil, err
	}
	err = s.Invoice.UpdateDB(ctx, invo)
	if err != nil {
		return nil, err
	}
	return invo.MakePDF()
}
