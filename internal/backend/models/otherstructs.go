package models

import (
	"encoding/xml"
)

type InvoicesDoc struct {
	XMLName  xml.Name  `xml:"InvoicesDoc"`
	Invoices []Invoice `xml:"Invoice"`
}

type Invoice struct {
	UID         string `xml:"uid,omitempty" json:"uid,omitempty"`
	Mark        int64  `xml:"mark,omitempty" json:"mark,omitempty"`
	CancelledBy int64  `xml:"cancelledByMark,omitempty" json:"cancelledByMark,omitempty"`

	InvoiceHeader  InvoiceHeader  `xml:"InvoiceHeader" json:"InvoiceHeader"`
	Issuer         Party          `xml:"Issuer" json:"Issuer"`
	Counterpart    Party          `xml:"Counterpart" json:"Counterpart"`
	InvoiceDetails []InvoiceLine  `xml:"InvoiceDetails>InvoiceRow" json:"InvoiceDetails"`
	InvoiceSummary InvoiceSummary `xml:"InvoiceSummary" json:"InvoiceSummary"`
}

type InvoiceHeader struct {
	Series               string `xml:"series" json:"series"`
	AA                   string `xml:"aa" json:"aa"`
	IssueDate            string `xml:"issueDate" json:"issueDate"`
	InvoiceType          string `xml:"invoiceType" json:"invoiceType"`
	VatPaymentSuspension bool   `xml:"vatPaymentSuspension,omitempty" json:"vatPaymentSuspension,omitempty"`
}

type Party struct {
	VATNumber  string `xml:"vatNumber" json:"vatNumber"`
	Country    string `xml:"country" json:"country"`
	Branch     int    `xml:"branch" json:"branch"`
	Name       string `xml:"name" json:"name"`
	Address    string `xml:"address" json:"address"`
	PostalCode string `xml:"postalCode,omitempty" json:"postalCode,omitempty"`
	City       string `xml:"city,omitempty" json:"city,omitempty"`
}

type InvoiceLine struct {
	LineNumber  int     `xml:"lineNumber" json:"lineNumber"`
	Quantity    float64 `xml:"quantity" json:"quantity"`
	UnitPrice   float64 `xml:"unitPrice" json:"unitPrice"`
	NetValue    float64 `xml:"netValue" json:"netValue"`
	VatCategory int     `xml:"vatCategory" json:"vatCategory"`
	VatAmount   float64 `xml:"vatAmount" json:"vatAmount"`
	Description string  `xml:"description" json:"description"`
}

type InvoiceSummary struct {
	TotalNetValue   float64 `xml:"totalNetValue" json:"totalNetValue"`
	TotalVatAmount  float64 `xml:"totalVatAmount" json:"totalVatAmount"`
	TotalGrossValue float64 `xml:"totalGrossValue" json:"totalGrossValue"`
}
