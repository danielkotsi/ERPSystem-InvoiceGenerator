package routes

import (
	"-invoice_manager/internal/backend/handlers"
	"net/http"
)

type Router struct {
	InvoiceHandler *handlers.InvoiceHandler
}

func (r *Router) Setup() http.Handler {
	mux := http.NewServeMux()

	// Get requests
	mux.HandleFunc("GET /", r.InvoiceHandler.GetHome)
	// mux.HandleFunc("GET /customers", r.InvoiceHandler.GetCustomers)
	// mux.HandleFunc("GET /products", r.InvoiceHandler.GetProducts)
	// mux.HandleFunc("GET /make_an_invoice", r.InvoiceHandler.GetMakeAnInvoicePage)
	//
	// //Post requests
	// mux.HandleFunc("POST /create-customer", r.InvoiceHandler.CreateCustomer)
	// mux.HandleFunc("POST /create-product", r.InvoiceHandler.CreateProduct)
	// mux.HandleFunc("POST /create-invoice", r.InvoiceHandler.CreateInvoice)
	return mux
}
