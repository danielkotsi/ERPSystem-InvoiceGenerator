package models

import (
	"-invoice_manager/internal/backend/invoice/payload"
)

type InvoiceHTMLinfo struct {
	User        payload.Company
	Invoiceinfo invoforminfo
}

type invoforminfo struct {
	Invoicetype              string
	Currency                 string
	Series                   string
	Aa                       string
	MovePurpose              string
	IncomeClassificationType string
	IncomeClassificationCat  string
	IsDeliveryNote           bool
	VatCategory              int
}
