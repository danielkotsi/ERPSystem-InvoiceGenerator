package mydata

import (
	"context"
	reposinterfaces "invoice_manager/internal/backend/invoice/reposInterfaces"
	"invoice_manager/internal/backend/invoice/types"
)

type MyDataRepo struct {
	client *Client
}

func NewMyDataRepo() *MyDataRepo {
	return &MyDataRepo{}
}

func (m *MyDataRepo) SendInvoice(ctx context.Context, invoice reposinterfaces.Invoice_type) (err error) {
	invo := invoice.GetInvoice()
	if types.InvoiceType(invo.InvoiceHeader.InvoiceType) == types.BuyingInvoiceType {
		return nil
	}
	//the variables passed here are later going to be retrived from the db based on the user which will be recognised from the auth middleware based on a cookie and a session, these values are probably going to be passed inside the context for easy acces from everywhere(among others) for now i will use my info nad keys by explicitely stating them
	m.client = NewMyDataClient("", "", "")
	//this is a new comment to check if we achieved what we wanted

	myDataResponse, err := m.client.SendInvoice(ctx, invo)
	if err != nil {
		return err
	}
	if myDataResponse.Response.StatusCode != "Success" {
		return &myDataResponse

	}

	invo.UID = myDataResponse.Response.InvoiceUID
	invo.QrURL = myDataResponse.Response.QRURL
	invo.MARK = myDataResponse.Response.InvoiceMARK

	return nil
}
