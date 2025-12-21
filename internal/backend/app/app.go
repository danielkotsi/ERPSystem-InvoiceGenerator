package app

import (
	"database/sql"
	"-invoice_manager/internal/backend/db/sqlite"
	"-invoice_manager/internal/backend/handlers"
	"-invoice_manager/internal/backend/middleware"
	"-invoice_manager/internal/backend/mydata"
	"-invoice_manager/internal/backend/routes"
	"-invoice_manager/internal/backend/services"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func New() (http.Handler, *sql.DB) {
	var exeDir string
	if os.Getenv("DEV") == "1" {
		exeDir, _ = filepath.Abs("../../")
	} else {
		exePath, _ := os.Executable()
		exeDir = filepath.Dir(exePath)
	}
	db := sqlite.NewDatabase(exeDir)
	templatesDir := filepath.Join(exeDir, "assets", "templates", "*.page.html")
	tmpl := template.Must(template.ParseGlob(templatesDir))

	invoiceRepo := sqlite.NewInvoiceRepo(db, exeDir)
	customersRepo := sqlite.NewCustomersRepo(db)
	productsRepo := sqlite.NewProductsRepo(db)
	myDataRepo := mydata.NewMyDataRepo()

	// Services
	invoice_service := services.NewInvoiceService(invoiceRepo, myDataRepo)
	customers_service := services.NewCustomersService(customersRepo)
	products_service := services.NewProductsService(productsRepo)
	htmlexcecuteservice := services.NewHTMLExcecutor(tmpl)

	// Handlers
	invoiceHandler := handlers.NewInvoiceHandler(invoice_service, htmlexcecuteservice)
	customersHandler := handlers.NewCustomersHandler(customers_service, htmlexcecuteservice)
	productsHandler := handlers.NewProductsHandler(products_service, htmlexcecuteservice)

	middleware := middleware.NewMiddleware()

	// Router
	router := &routes.Router{
		InvoiceHandler:   invoiceHandler,
		CustomersHandler: customersHandler,
		ProductsHandler:  productsHandler,
		Middleware:       middleware,
	}

	return router.Setup(exeDir), db
}
