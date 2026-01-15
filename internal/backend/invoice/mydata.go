package invoice

import (
	"context"
	"-invoice_manager/internal/backend/invoice/payload"
)

type MyData_repo interface {
	SendInvoice(ctx context.Context, invoice *payload.Invoice) (err error)
}
