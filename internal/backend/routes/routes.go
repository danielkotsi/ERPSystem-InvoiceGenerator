package routes

import (
	"-invoice_manager/internal/backend/customer"
	"-invoice_manager/internal/backend/invoice"
	"-invoice_manager/internal/backend/middleware"
	"-invoice_manager/internal/backend/product"
	"net/http"
	"path/filepath"
)

type Router struct {
	InvoiceHandler   *invoice.InvoiceHandler
	CustomersHandler *customer.CustomersHandler
	ProductsHandler  *product.ProductsHandler
	Middleware       *middleware.Middleware
}

func NewRouter(invoiceHandler *invoice.InvoiceHandler, customersHandler *customer.CustomersHandler, productsHandler *product.ProductsHandler, middleware *middleware.Middleware) *Router {
	return &Router{
		InvoiceHandler:   invoiceHandler,
		CustomersHandler: customersHandler,
		ProductsHandler:  productsHandler,
		Middleware:       middleware,
	}
}
func (r *Router) Setup(abspath string) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(abspath, "static")))))
	mux.HandleFunc("GET /", r.InvoiceHandler.GetHome)
	mux.HandleFunc("GET /customers", r.CustomersHandler.GetCustomers)
	mux.HandleFunc("GET /customers/byid/", r.CustomersHandler.GetCustomerById)
	mux.HandleFunc("GET /suggestions/customers", r.CustomersHandler.GetCustomerSuggestions)
	mux.HandleFunc("GET /suggestions/full/customer", r.CustomersHandler.GetFullCustomer)
	mux.HandleFunc("GET /suggestions/branchcompanies", r.CustomersHandler.GetBranchCompaniesSuggestions)
	mux.HandleFunc("GET /suggestions/full/branchcompany", r.CustomersHandler.GetFullBranchCompany)
	mux.HandleFunc("POST /create/branchcompany", r.CustomersHandler.CreateBranchCompany)
	mux.HandleFunc("POST /create/customer", r.CustomersHandler.CreateCustomer)

	mux.HandleFunc("GET /products", r.ProductsHandler.GetProducts)
	mux.HandleFunc("GET /suggestions/products", r.ProductsHandler.GetProductSuggestions)
	mux.HandleFunc("GET /suggestions/full/product", r.ProductsHandler.GetProductById)
	mux.HandleFunc("POST /create/product", r.ProductsHandler.CreateProduct)

	mux.HandleFunc("GET /makeaninvoice", r.InvoiceHandler.GetMakeInvoicePage)
	mux.HandleFunc("POST /invoice/create", r.InvoiceHandler.CreateInvoice)
	return r.Middleware.Handler(mux)
}
