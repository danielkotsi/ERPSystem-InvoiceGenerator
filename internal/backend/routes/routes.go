package routes

import (
	"-invoice_manager/internal/backend/handlers"
	"-invoice_manager/internal/backend/middleware"
	"net/http"
)

type Router struct {
	InvoiceHandler   *handlers.InvoiceHandler
	CustomersHandler *handlers.CustomersHandler
	ProductsHandler  *handlers.ProductsHandler
	Middleware       *middleware.Middleware
}

func (r *Router) Setup() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../static/"))))
	// Get requests
	mux.HandleFunc("GET /", r.InvoiceHandler.GetHome)
	mux.HandleFunc("GET /customers", r.CustomersHandler.GetCustomers)
	mux.HandleFunc("GET /suggestions/customers", r.CustomersHandler.GetCustomerSuggestions)
	mux.HandleFunc("POST /customers/create", r.CustomersHandler.CreateCustomer)

	mux.HandleFunc("GET /products", r.ProductsHandler.GetProducts)
	mux.HandleFunc("GET /suggestions/products", r.ProductsHandler.GetProductSuggestions)
	mux.HandleFunc("POST /products/create", r.ProductsHandler.CreateProduct)

	// mux.HandleFunc("GET /invoice/list", r.ProductsHandler.GetInvoices)
	mux.HandleFunc("GET /makeaninvoice", r.InvoiceHandler.GetMakeInvoicePage)
	mux.HandleFunc("POST /invoice/create", r.InvoiceHandler.CreateInvoice)
	// mux.HandleFunc("GET /make_an_invoice", r.InvoiceHandler.GetMakeAnInvoicePage)
	//
	// //Post requests
	// mux.HandleFunc("POST /create-customer", r.InvoiceHandler.CreateCustomer)
	// mux.HandleFunc("POST /create-product", r.InvoiceHandler.CreateProduct)
	// mux.HandleFunc("POST /create-invoice", r.InvoiceHandler.CreateInvoice)
	return r.Middleware.Handler(mux)
}
