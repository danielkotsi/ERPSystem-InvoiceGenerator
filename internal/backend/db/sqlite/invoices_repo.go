package sqlite

import (
	"bytes"
	"context"
	"database/sql"
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/utils"
	"html/template"
	"log"
)

type InvoiceRepo struct {
	DB *sql.DB
}

func NewInvoiceRepo(db *sql.DB) *InvoiceRepo {
	return &InvoiceRepo{DB: db}
}

func (r *InvoiceRepo) DesignInvoice(ctx context.Context, invo models.Invoice) (pdf []byte, err error) {

	tmpl, err := template.ParseFiles("../../assets/templates/invoice.page.html")
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, map[string]models.Invoice{"Invoice": invo})
	if err != nil {
		log.Println(err)
	}
	pdf, err = utils.HTMLtoPDF(buf.String())
	if err != nil {
		return nil, err
	}

	return pdf, nil
}
