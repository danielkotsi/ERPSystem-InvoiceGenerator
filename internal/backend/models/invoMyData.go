package models

import "encoding/xml"

type InvoicePayload struct {
	XMLName xml.Name `xml:"InvoicePayload"`
	Invoice Invoice  `json:"invoice" form:"invoice" xml:"invoice"`
}

//	type Invoice struct {
//		Seller         Company               `json:"issuer" form:"seller" xml:"issuer"`
//		Byer           Company               `json:"counterpart" form:"buyer" xml:"counterpart"`
//		InvoiceHeader  InvoiceHeader         `json:"invoiceHeader" form:"invoiceHeader" xml:"invoiceHeader"`
//		PaymentMethods []PaymentMethodDetail `json:"paymentMethods,omitempty" form:"paymentMethods" xml:"paymentMethods"`
//		InvoiceDetails []*InvoiceRow         `json:"invoiceDetails" form:"invoiceDetails" xml:"invoiceDetails"`
//		InvoiceSummary InvoiceSummary        `json:"invoiceSummary" form:"invoiceSummary" xml:"invoiceSummary"`
//	}
type Company struct {
	VatNumber  string       `json:"vatNumber" form:"vatNumber" xml:"vatNumber"`
	Country    string       `json:"country" form:"country" xml:"country"`
	Branch     int          `json:"branch" form:"branch" xml:"branch"`
	EntityType int          `json:"entityType" form:"entityType" xml:"entityType"`
	Name       string       `json:"name" form:"name" xml:"name"`
	Address    *AddressType `json:"address,omitempty" form:"address" xml:"address"`
}

// type InvoiceHeader struct {
// 	Series      string `json:"series" form:"series" xml:"series"`
// 	Aa          string `json:"aa" form:"aa" xml:"aa"`
// 	IssueDate   string `json:"issueDate" form:"issueDate" xml:"issueDate"`
// 	InvoiceType string `json:"invoiceType" form:"invoiceType" xml:"invoiceType"`
// }

type PaymentMethodDetail struct {
	Type   int     `json:"type" form:"type" xml:"type"`
	Amount float64 `json:"amount" form:"amount" xml:"amount"`
	TID    string  `json:"tid,omitempty" form:"tid" xml:"tid"`
}
type InvoiceRow struct {
	LineNumber  int     `json:"lineNumber" form:"lineNumber" xml:"lineNumber"`
	RecType     int     `json:"recType,omitempty" form:"recType" xml:"recType"`
	Quantity    float64 `json:"quantity,omitempty" form:"quantity" xml:"quantity"`
	UnitPrice   float64 `json:"unitPrice,omitempty" form:"unitPrice" xml:"unitPrice"`
	NetValue    float64 `json:"netValue" form:"netValue" xml:"netValue"`
	VatCategory int     `json:"vatCategory" form:"vatCategory" xml:"vatCategory"`
	VatAmount   float64 `json:"vatAmount" form:"vatAmount" xml:"vatAmount"`
	Description string  `json:"description" form:"description" xml:"description"`
}

//	type InvoiceSummary struct {
//		TotalNetValue        float64              `json:"totalNetValue" form:"totalNetValue" xml:"totalNetValue"`
//		TotalVatAmount       float64              `json:"totalVatAmount" form:"totalVatAmount" xml:"totalVatAmount"`
//		TotalWithVat         float64              `json:"totalWithVat" form:"totalWithVat" xml:"totalWithVat"`
//		IncomeClassification []ClassificationItem `json:"incomeClassification,omitempty" form:"incomeClassification" xml:"incomeClassification"`
//	}
type ClassificationItem struct {
	ClassificationType     string  `json:"classificationType" form:"classificationType" xml:"classificationType"`
	ClassificationCategory string  `json:"classificationCategory" form:"classificationCategory" xml:"classificationCategory"`
	Amount                 float64 `json:"amount" form:"amount" xml:"amount"`
}
