package models

import (
	"net/http"
)

type Client struct {
	BaseURL         string
	UserID          string
	SubscriptionKey string
	HTTPClient      *http.Client
}
