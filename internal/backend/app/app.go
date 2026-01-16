package app

import (
	"database/sql"
	"-invoice_manager/internal/backend/customer"
	"-invoice_manager/internal/backend/invoice"
	"-invoice_manager/internal/backend/invoice/adapter"
	"-invoice_manager/internal/backend/middleware"
	"-invoice_manager/internal/backend/product"
	"-invoice_manager/internal/backend/routes"
	"-invoice_manager/internal/backend/services"
	"-invoice_manager/internal/infrastructure/db/sqlite"
	"-invoice_manager/internal/infrastructure/mydata"
	"-invoice_manager/internal/utils"
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
	logo := utils.Imageto64(exeDir)
	db := sqlite.NewDatabase(exeDir)
	templatesDir := filepath.Join(exeDir, "assets", "templates", "*.page.html")
	tmpl := template.Must(template.ParseGlob(templatesDir))

	invoiceRepo := sqlite.NewInvoiceRepo(db, exeDir, logo)
	customersRepo := sqlite.NewCustomersRepo(db)
	productsRepo := sqlite.NewProductsRepo(db)
	myDataRepo := mydata.NewMyDataRepo()

	// Services
	invoice_service := invoice.NewInvoiceService(invoiceRepo, myDataRepo)
	customers_service := customer.NewCustomersService(customersRepo)
	products_service := product.NewProductsService(productsRepo)
	htmlexcecuteservice := services.NewHTMLExcecutor(tmpl)

	//the adapter for the Invoices
	invoiceAdapter := adapter.NewInvoiceParser()
	// Handlers
	invoiceHandler := invoice.NewInvoiceHandler(invoice_service, htmlexcecuteservice, invoiceAdapter)
	customersHandler := customer.NewCustomersHandler(customers_service, htmlexcecuteservice)
	productsHandler := product.NewProductsHandler(products_service, htmlexcecuteservice)

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
