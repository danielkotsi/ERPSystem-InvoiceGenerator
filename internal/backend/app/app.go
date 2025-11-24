package app

import (
	"database/sql"
	"-invoice_manager/internal/backend/db/sqlite"
	"-invoice_manager/internal/backend/handlers"
	"-invoice_manager/internal/backend/middleware"
	"-invoice_manager/internal/backend/routes"
	"-invoice_manager/internal/backend/services"
	"-invoice_manager/internal/utils"
	"html/template"
	"net/http"
)

func New() (http.Handler, *sql.DB) {
	CONFIG := utils.DecodeConf()
	db := sqlite.NewDatabase(&CONFIG)
	tmpl := template.Must(template.ParseGlob("../../assets/templates/*.page.html"))
	// Repos
	invoiceRepo := sqlite.NewInvoiceRepo(db)

	// Services
	invoice_service := services.NewInvoiceService(invoiceRepo)
	htmlexcecuteservice := services.NewHTMLExcecutor(tmpl)

	// Handlers
	invoiceHandler := handlers.NewInvoiceHandler(invoice_service, htmlexcecuteservice)

	middleware := middleware.NewMiddleware(&CONFIG)

	// Router
	router := &routes.Router{
		InvoiceHandler: invoiceHandler,
		Middleware:     middleware,
	}

	return router.Setup(), db
}
