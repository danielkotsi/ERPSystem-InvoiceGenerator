package invoice

import (
	"context"
	"-invoice_manager/internal/backend/invoice/models"
	"-invoice_manager/internal/backend/invoice/reposInterfaces"
	"-invoice_manager/internal/backend/invoice/types"
	"net/http"
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

func (s *InvoiceService) GetInvoiceInfo(ctx context.Context, invoicetype types.InvoiceType) (invoiceinfo models.InvoiceHTMLinfo, invoiceHTML string, err error) {

	invoiceinfo, err = s.Invoice.GetInvoiceInfo(ctx, invoicetype)
	if err != nil {
		return invoiceinfo, string(types.InvoiceHTML(string(invoicetype))), err
	}
	return invoiceinfo, string(types.InvoiceHTML(string(invoicetype))), nil
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, invo reposinterfaces.Invoice_type) (pdf []byte, err error) {
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
