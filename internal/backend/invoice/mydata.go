package invoice

import (
	"context"
	"invoice_manager/internal/backend/invoice/reposInterfaces"
)

type MyData_repo interface {
	SendInvoice(ctx context.Context, invoice reposinterfaces.Invoice_type) (err error)
}
