package models

import "encoding/xml"

type InvoicePayload struct {
	XMLName  xml.Name  `xml:"http://www.aade.gr/myDATA/invoice/v1.0 InvoicesDoc"`
	Invoices []Invoice `form:"invoice" xml:"invoice"`
}

type Invoice struct {
	Seller         Company        `json:"issuer" form:"seller" xml:"issuer"`
	Byer           Company        `json:"counterpart" form:"buyer" xml:"counterpart"`
	InvoiceHeader  InvoiceHeader  `json:"invoiceHeader" form:"invoiceHeader" xml:"invoiceHeader"`
	PaymentMethods PaymentMethods `json:"paymentMethods,omitempty" form:"paymentMethods" xml:"paymentMethods"`
	InvoiceDetails []*InvoiceRow  `json:"invoiceDetails" form:"invoiceDetails" xml:"invoiceDetails"`
	InvoiceSummary InvoiceSummary `json:"invoiceSummary" form:"invoiceSummary" xml:"invoiceSummary"`
}

type Company struct {
	VatNumber string       `json:"vatNumber" form:"vatNumber" xml:"vatNumber"`
	Country   string       `json:"country" form:"country" xml:"country"`
	Branch    int          `json:"branch" form:"branch" xml:"branch"`
	Name      string       `json:"name" form:"name" xml:"name"`
	Address   *AddressType `json:"address,omitempty" form:"address" xml:"address"`
}

type AddressType struct {
	Street     *string `json:"street,omitempty" form:"street" xml:"street"`
	Number     *string `json:"number,omitempty" form:"number" xml:"number"`
	PostalCode *string `json:"postalCode,omitempty" form:"postalCode" xml:"postalCode"`
	City       *string `json:"city,omitempty" form:"city" xml:"city"`
}

type InvoiceHeader struct {
	Series      string `json:"series" form:"series" xml:"series"`
	Aa          string `json:"aa" form:"aa" xml:"aa"`
	IssueDate   string `json:"issueDate" form:"issueDate" xml:"issueDate"`
	InvoiceType string `json:"invoiceType" form:"invoiceType" xml:"invoiceType"`
	Currency    string `json:"currency" form:"currency" xml:"currency"`
}

type PaymentMethods struct {
	Details []PaymentMethodDetail `json:"paymentdatails" form:"paymentdetails" xml:"paymentMethodDetails"`
}
type PaymentMethodDetail struct {
	Type   int     `json:"type" form:"type" xml:"type"`
	Amount float64 `json:"amount" form:"amount" xml:"amount"`
}
type InvoiceRow struct {
	LineNumber int `json:"lineNumber" form:"lineNumber" xml:"lineNumber"`
	// RecType     int     `json:"recType,omitempty" form:"recType" xml:"recType"`
	Quantity float64 `json:"quantity,omitempty" form:"quantity" xml:"quantity"`
	//this is not tested on the xml therefore, when a bug is encountered this needs to be checked
	UnitPrice   float64 `json:"unitPrice,omitempty" form:"unitPrice" xml:"unitPrice"`
	NetValue    float64 `json:"netValue" form:"netValue" xml:"netValue"`
	VatCategory int     `json:"vatCategory" form:"vatCategory" xml:"vatCategory"`
	VatAmount   float64 `json:"vatAmount" form:"vatAmount" xml:"vatAmount"`
	// Description string  `json:"description" form:"description" xml:"description"`
	IncomeClassification ClassificationItem `json:"incomeClassification" form:"incomeClassification" xml:"incomeClassification"`
}

type InvoiceSummary struct {
	TotalNetValue         float64              `json:"totalNetValue" form:"totalNetValue" xml:"totalNetValue"`
	TotalVatAmount        float64              `json:"totalVatAmount" form:"totalVatAmount" xml:"totalVatAmount"`
	TotalWithheldAmount   float64              `json:"totalWithheldAmount" form:"totalWithheldAmount" xml:"totalWithheldAmount"`
	TotalFeesAmount       float64              `json:"totalFeesAmount" form:"totalFeesAmount" xml:"totalFeesAmount"`
	TotalStampDutyAmount  float64              `json:"totalStampDutyAmount" form:"totalStampDutyAmount" xml:"totalStampDutyAmount"`
	TotalOtherTaxesAmount float64              `json:"totalOtherTaxesAmount" form:"totalOtherTaxesAmount" xml:"totalOtherTaxesAmount"`
	TotalDeductionsAmount float64              `json:"totalDeductionsAmount" form:"totalDeductionsAmount" xml:"totalDeductionsAmount"`
	TotalGrossValue       float64              `json:"totalGrossValue" form:"totalGrossValue" xml:"totalGrossValue"`
	IncomeClassification  []ClassificationItem `json:"incomeClassification,omitempty" form:"incomeClassification" xml:"incomeClassification"`
}

type ClassificationItem struct {
	ClassificationType     string  `json:"classificationType" form:"classificationType" xml:"https://www.aade.gr/myDATA/incomeClassificaton/v1.0 classificationType"`
	ClassificationCategory string  `json:"classificationCategory" form:"classificationCategory" xml:"https://www.aade.gr/myDATA/incomeClassificaton/v1.0 classificationCategory"`
	Amount                 float64 `json:"amount" form:"amount" xml:"https://www.aade.gr/myDATA/incomeClassificaton/v1.0 amount"`
}
