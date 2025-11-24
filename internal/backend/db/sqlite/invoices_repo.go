package sqlite

import (
	"bytes"
	"database/sql"
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

func (r *InvoiceRepo) DesignInvoice() error {

	tmpl, err := template.ParseFiles("../../assets/templates/home.page.html")
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		log.Println(err)
	}
	utils.HTMLtoPDF(buf.String(), "../../../../config/thenewpdffile.pdf")

	return nil
}
