package mydata

import (
	"bytes"
	"context"
	"-invoice_manager/internal/backend/models"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	BaseURL         string
	UserID          string
	SubscriptionKey string
	HTTPClient      *http.Client
}

func NewMyDataClient(base, userID, subscriptionKey string) *Client {
	//this is going to change afterwards cause i dont know whoose keys are going to be used
	// base = os.Getenv("BASE_URL")
	// subscriptionKey = os.Getenv("API_KEY")
	// userID = os.Getenv("USER")
	return &Client{
		BaseURL:         base,
		UserID:          userID,
		SubscriptionKey: subscriptionKey,
		HTTPClient:      &http.Client{Timeout: 15 * time.Second},
	}

}

func (c *Client) SendInvoice(ctx context.Context, invoice *models.Invoice) (models.ResponseDoc, error) {

	var invoicePayload models.InvoicePayload
	invoicePayload.Invoices = append(invoicePayload.Invoices, *invoice)
	invo, err := xml.MarshalIndent(invoicePayload, "", "  ")
	if err != nil {
		return models.ResponseDoc{}, err
	}

	fmt.Println("this is the invoice after the marshalling\n", string(invo))

	response, err := c.DoRequest(invo)
	if err != nil {
		return models.ResponseDoc{}, err
	}

	var myDataResponse models.ResponseDoc
	err = xml.Unmarshal(response, &myDataResponse)
	if err != nil {
		return models.ResponseDoc{}, err
	}
	return myDataResponse, nil
}

func ImportXML() (invo []byte, err error) {
	file, err := os.Open("../../exampleInvoice.xml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	invo, err = io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return invo, nil
}
func (c *Client) DoRequest(invo []byte) (response []byte, err error) {
	url := c.BaseURL + "SendInvoices"
	fmt.Println(url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(invo))
	if err != nil {
		log.Println("this is the request error", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("aade-user-id", c.UserID)
	req.Header.Set("ocp-apim-subscription-key", c.SubscriptionKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Println("this is the error from my data", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Println("this is the response", string(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return body, fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}

	return body, nil
}
