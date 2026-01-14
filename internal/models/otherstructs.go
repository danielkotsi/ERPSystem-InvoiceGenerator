package models

type CustomerById struct {
	Customer        Company
	BranchCompanies []BranchCompany
}

type InvoiceHTMLinfo struct {
	User        Company
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
