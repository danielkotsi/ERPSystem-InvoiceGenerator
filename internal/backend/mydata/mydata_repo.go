package mydata

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"log"
)

type MyDataRepo struct {
	client *Client
}

func NewMyDataRepo() *MyDataRepo {
	return &MyDataRepo{}
}

func (m *MyDataRepo) SendInvoice(ctx context.Context, invoicePayload *models.InvoicePayload) (completeinvo []byte, err error) {
	//the variables passed here are later going to be retrived from the db based on the user which will be recognised from the auth middleware based on a cookie and a session, these values are probably going to be passed inside the context for easy acces from everywhere(among others) for now i will use my info nad keys by explicitely stating them
	m.client = NewMyDataClient("", "", "")
	//this is a new comment to check if we achieved what we wanted

	completeinvo, err = m.client.SendInvoice(ctx, invoicePayload)
	if err != nil {
		return completeinvo, err
	}

	log.Println(string(completeinvo))
	return completeinvo, nil
}
