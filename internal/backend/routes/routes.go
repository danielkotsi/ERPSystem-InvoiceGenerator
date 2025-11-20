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

	// Sessions
	mux.HandleFunc("GET /", r.InvoiceHandler.GetHome)

	// Wrap with middleware
	return mux
}
