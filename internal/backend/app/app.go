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
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func New() (http.Handler, *sql.DB) {
	exeDir, logo, db, tmpl := NewAppConsts()

	invoiceRepo, customersRepo, productsRepo, myDataRepo := NewRepos(db, exeDir, logo)

	invoice_service, customers_service, products_service, htmlexcecuteservice := NewServices(invoiceRepo, customersRepo, productsRepo, myDataRepo, tmpl)

	invoiceAdapter := adapter.NewInvoiceParser(logo, exeDir)

	invoiceHandler, customersHandler, productsHandler := NewHandlers(invoice_service, customers_service, products_service, htmlexcecuteservice, invoiceAdapter)

	middleware := middleware.NewMiddleware()

	router := routes.NewRouter(invoiceHandler, customersHandler, productsHandler, middleware)

	return router.Setup(exeDir), db
}

func NewRepos(db *sql.DB, exeDir string, logo string) (invoicerepo *sqlite.InvoiceRepo, customersRepo *sqlite.CustomersRepo, productsRepo *sqlite.ProductsRepo, myDataRepo *mydata.MyDataRepo) {

	invoicerepo, err := sqlite.NewInvoiceRepo(db, exeDir, logo)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	customersRepo, err = sqlite.NewCustomersRepo(db)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	productsRepo = sqlite.NewProductsRepo(db)
	myDataRepo = mydata.NewMyDataRepo()

	return invoicerepo,
		customersRepo,
		productsRepo,
		myDataRepo
}

func NewAppConsts() (exeDir, logo string, db *sql.DB, tmpl *template.Template) {
	if os.Getenv("DEV") == "1" {
		exeDir, _ = filepath.Abs("../../")
	} else {
		exePath, _ := os.Executable()
		exeDir = filepath.Dir(exePath)
	}
	logo = filepath.Join(exeDir, "static", "images", "logo.png")
	db = sqlite.NewDatabase(exeDir)
	templatesDir := filepath.Join(exeDir, "assets", "templates", "*.page.html")
	tmpl = template.Must(template.ParseGlob(templatesDir))
	return exeDir, logo, db, tmpl
}
func NewServices(invoicerepo *sqlite.InvoiceRepo, customersRepo *sqlite.CustomersRepo, productsRepo *sqlite.ProductsRepo, myDataRepo *mydata.MyDataRepo, tmpl *template.Template) (invoice_service *invoice.InvoiceService, customers_service *customer.CustomersService, products_service *product.ProductsService, htmlexcecuteservice *services.Excecutor) {

	invoice_service = invoice.NewInvoiceService(invoicerepo, myDataRepo)
	customers_service = customer.NewCustomersService(customersRepo)
	products_service = product.NewProductsService(productsRepo)
	htmlexcecuteservice = services.NewHTMLExcecutor(tmpl)

	return invoice_service,
		customers_service,
		products_service,
		htmlexcecuteservice
}

func NewHandlers(invoice_service *invoice.InvoiceService, customers_service *customer.CustomersService, products_service *product.ProductsService, htmlexcecuteservice *services.Excecutor, invoiceAdapter *adapter.InvoiceParser) (invoiceHandler *invoice.InvoiceHandler, customersHandler *customer.CustomersHandler, productsHandler *product.ProductsHandler) {
	invoiceHandler = invoice.NewInvoiceHandler(invoice_service, htmlexcecuteservice, invoiceAdapter)
	customersHandler = customer.NewCustomersHandler(customers_service, htmlexcecuteservice)
	productsHandler = product.NewProductsHandler(products_service, htmlexcecuteservice)
	return invoiceHandler, customersHandler, productsHandler
}
