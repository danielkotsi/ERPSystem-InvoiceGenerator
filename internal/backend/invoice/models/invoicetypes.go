package models

import (
	"-invoice_manager/internal/backend/invoice/payload"
)

type SellingInvoice struct {
	Payload *payload.InvoicePayload
}

type BuyingInvoice struct {
	Payload *payload.InvoicePayload
}

type Reciept struct {
	Payload *payload.InvoicePayload
}

type DeliveryNote struct {
	Payload *payload.InvoicePayload
}
