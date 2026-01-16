package types

type InvoiceType string

const (
	SellingInvoiceType      InvoiceType = "1.1"
	BuyingInvoiceType       InvoiceType = "13.1"
	DeliveryNoteInvoiceType InvoiceType = "9.3"
	RecieptInvoiceType      InvoiceType = "8.1"
)

type InvoiceHTML string

const (
	SellingInvoiceHTML      InvoiceHTML = "1.1"
	BuyingInvoiceHTML       InvoiceHTML = "13.1"
	DeliveryNoteInvoiceHTML InvoiceHTML = "9.3"
	RecieptInvoiceHTML      InvoiceHTML = "8.1"
)
