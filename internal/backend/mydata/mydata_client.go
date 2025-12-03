package mydata

import (
	"bytes"
	"context"
	"-invoice_manager/internal/backend/models"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
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

func (c *Client) SendInvoice(ctx context.Context, invoicepayload *models.InvoicePayload) ([]byte, error) {

	invo, err := xml.MarshalIndent(invoicepayload, "", "  ")
	if err != nil {
		return nil, err
	}

	completeinvo, err := c.DoRequest(invo)
	if err != nil {
		return nil, err
	}
	return completeinvo, nil
}

func (c *Client) DoRequest(invo []byte) (completeinvo []byte, err error) {
	url := c.BaseURL + "/SendInvoices"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(invo))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("aade-user-id", c.UserID)
	req.Header.Set("ocp-apim-subscription-key", c.SubscriptionKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return body, errors.New(fmt.Sprintf("HTTP Error: %d", resp.StatusCode))
	}

	return body, nil
}
