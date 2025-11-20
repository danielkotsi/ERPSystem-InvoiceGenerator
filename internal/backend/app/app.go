package app

import (
	"database/sql"
	"-invoice_manager/internal/backend/db/sqlite"
	"-invoice_manager/internal/backend/handlers"
	"-invoice_manager/internal/backend/routes"
	"-invoice_manager/internal/backend/services"
	"-invoice_manager/internal/utils"
	"net/http"
)

func New() (http.Handler, *sql.DB) {
	CONFIG := utils.DecodeConf()
	db := sqlite.NewDatabase(&CONFIG)

	// Repos
	invoiceRepo := sqlite.NewInvoiceRepo(db)

	// Services
	invoice_service := services.NewInvoiceService(invoiceRepo)

	// Handlers
	invoiceHandler := &handlers.InvoiceHandler{InvoiceService: invoice_service}

	// Router
	router := &routes.Router{
		InvoiceHandler: invoiceHandler,
	}

	return router.Setup(), db
}
