package mydata

import (
	"context"
	"-invoice_manager/internal/backend/models"
	"net/http"
	"time"
)

type MyDataRepo struct {
	client *models.Client
}

func NewMyDataRepo() *MyDataRepo {
	return &MyDataRepo{}
}

func (m *MyDataRepo) NewMyDataClient(base, userID, subscriptionKey string) {
	m.client = &models.Client{
		BaseURL:         base,
		UserID:          userID,
		SubscriptionKey: subscriptionKey,
		HTTPClient:      &http.Client{Timeout: 15 * time.Second},
	}

}

func (m *MyDataRepo) SendInvoice(ctx context.Context, invoicePayload *models.InvoicePayload) error {
	//the variables passed here are later going to be retrived from the db based on the user which will be recognised from the auth middleware based on a cookie and a session, these values are probably going to be passed inside the context for easy acces from everywhere(among others) for now i will use my info nad keys by explicitely stating them
	m.NewMyDataClient("", "", "")
	return nil
}
