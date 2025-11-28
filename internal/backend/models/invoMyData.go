package models

type InvoicePayload struct {
	Invoice Invoice `json:"invoice" form:"invoice"`
}

type Invoice struct {
	Seller         Company               `json:"issuer" form:"seller"`
	Byer           Company               `json:"counterpart" form:"buyer"`
	InvoiceHeader  InvoiceHeader         `json:"invoiceHeader" form:"invoiceHeader"`
	PaymentMethods []PaymentMethodDetail `json:"paymentMethods,omitempty" form:"paymentMethods"`
	InvoiceDetails []InvoiceRow          `json:"invoiceDetails" form:"invoiceDetails"`
	InvoiceSummary InvoiceSummary        `json:"invoiceSummary" form:"invoiceSummary"`
}

type Company struct {
	VatNumber  string       `json:"vatNumber" form:"vatNumber"`
	Country    string       `json:"country" form:"country"`
	Branch     int          `json:"branch" form:"branch"`
	EntityType int          `json:"entityType" form:"entityType"`
	Name       string       `json:"name" form:"name"`
	Address    *AddressType `json:"address,omitempty" form:"address"`
}

type InvoiceHeader struct {
	Series      string `json:"series" form:"series"`
	Aa          string `json:"aa" form:"aa"`
	IssueDate   string `json:"issueDate" form:"issueDate"`
	InvoiceType string `json:"invoiceType" form:"invoiceType"`
}

type PaymentMethodDetail struct {
	Type   int     `json:"type" form:"type"`
	Amount float64 `json:"amount" form:"amount"`
	TID    string  `json:"tid,omitempty" form:"tid"`
}
type InvoiceRow struct {
	LineNumber  int     `json:"lineNumber" form:"lineNumber"`
	RecType     int     `json:"recType,omitempty" form:"recType"`
	Quantity    float64 `json:"quantity,omitempty" form:"quantity"`
	UnitPrice   float64 `json:"unitPrice,omitempty" form:"unitPrice"`
	NetValue    float64 `json:"netValue" form:"netValue"`
	VatCategory int     `json:"vatCategory" form:"vatCategory"`
	VatAmount   float64 `json:"vatAmount" form:"vatAmount"`
}

type InvoiceSummary struct {
	TotalNetValue        float64              `json:"totalNetValue" form:"totalNetValue"`
	TotalVatAmount       float64              `json:"totalVatAmount" form:"totalVatAmount"`
	TotalWithVat         float64              `json:"totalWithVat" form:"totalWithVat"`
	IncomeClassification []ClassificationItem `json:"incomeClassification,omitempty" form:"incomeClassification"`
}

type ClassificationItem struct {
	ClassificationType     string  `json:"classificationType" form:"classificationType"`
	ClassificationCategory string  `json:"classificationCategory" form:"classificationCategory"`
	Amount                 float64 `json:"amount" form:"amount"`
}
