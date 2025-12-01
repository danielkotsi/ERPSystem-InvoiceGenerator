package mydata

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL         string
	UserID          string
	SubscriptionKey string
	HTTPClient      *http.Client
}

func NewMyDataClient(base, userID, subscriptionKey string) *Client {
	return &Client{
		BaseURL:         base,
		UserID:          userID,
		SubscriptionKey: subscriptionKey,
		HTTPClient:      &http.Client{Timeout: 15 * time.Second},
	}

}

func (c *Client) SendInvoice() {
}
