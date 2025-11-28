package models

type InvoicePayload struct {
	Invoice Invoice `json:"invoice"`
}

type Invoice struct {
	Seller         Company               `json:"issuer"`
	Byer           Company               `json:"counterpart"`
	InvoiceHeader  InvoiceHeader         `json:"invoiceHeader"`
	PaymentMethods []PaymentMethodDetail `json:"paymentMethods,omitempty"`
	InvoiceDetails []InvoiceRow          `json:"invoiceDetails"`
	InvoiceSummary InvoiceSummary        `json:"invoiceSummary"`
}

type Company struct {
	VatNumber  string       `json:"vatNumber"`
	Country    string       `json:"country"`
	Branch     int          `json:"branch"`
	EntityType int          `json:"entityType"`
	Name       string       `json:"name"`
	Address    *AddressType `json:"address,omitempty"`
}

type AddressType struct {
	Street     *string `json:"street,omitempty"`
	Number     *string `json:"number,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	City       *string `json:"city,omitempty"`
}

type InvoiceHeader struct {
	Series      string `json:"series"`
	Aa          string `json:"aa"`
	IssueDate   string `json:"issueDate"`   // format: YYYY-MM-DD
	InvoiceType string `json:"invoiceType"` // ex: "1.1"
}

type PaymentMethodDetail struct {
	Type   int     `json:"type"` // 1=Cash, 2=Bank, 3=Card, 4=POS
	Amount float64 `json:"amount"`
	TID    string  `json:"tid,omitempty"` // optional POS terminal ID
}

type InvoiceRow struct {
	LineNumber  int     `json:"lineNumber"`
	RecType     int     `json:"recType,omitempty"` // default 3 for normal row
	Quantity    float64 `json:"quantity,omitempty"`
	UnitPrice   float64 `json:"unitPrice,omitempty"`
	NetValue    float64 `json:"netValue"`
	VatCategory int     `json:"vatCategory"`
	VatAmount   float64 `json:"vatAmount"`
}

type InvoiceSummary struct {
	TotalNetValue        float64              `json:"totalNetValue"`
	TotalVatAmount       float64              `json:"totalVatAmount"`
	TotalWithVat         float64              `json:"totalWithVat"`
	IncomeClassification []ClassificationItem `json:"incomeClassification,omitempty"`
}

type ClassificationItem struct {
	ClassificationType     string  `json:"classificationType"`     // ex: "E3_106"
	ClassificationCategory string  `json:"classificationCategory"` // ex: "category1_3"
	Amount                 float64 `json:"amount"`
}
