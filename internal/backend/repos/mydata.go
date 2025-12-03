package repository

import (
	"context"
	"-invoice_manager/internal/backend/models"
)

type MyData_repo interface {
	SendInvoice(ctx context.Context, invoice *models.InvoicePayload) (completeinvo []byte, err error)
}
