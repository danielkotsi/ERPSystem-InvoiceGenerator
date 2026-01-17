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
	SellingInvoiceHTML      InvoiceHTML = "create_selling_invoice.page.html"
	BuyingInvoiceHTML       InvoiceHTML = "create_buying_invoice.page.html"
	DeliveryNoteInvoiceHTML InvoiceHTML = "create_deliverynote_invoice.page.html"
	RecieptInvoiceHTML      InvoiceHTML = "create_reciept_invoice.page.html"
)

func (r InvoiceType) HTMLTemplate() InvoiceHTML {
	switch r {
	case SellingInvoiceType:
		return SellingInvoiceHTML
	case BuyingInvoiceType:
		return BuyingInvoiceHTML
	case DeliveryNoteInvoiceType:
		return DeliveryNoteInvoiceHTML
	case RecieptInvoiceType:
		return RecieptInvoiceHTML
	default:
		return ""
	}
}
