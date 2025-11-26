package handlers

import (
	"-invoice_manager/internal/backend/services"
	"net/http"
)

type InvoiceHandler struct {
	InvoiceService *services.InvoiceService
	Excecutor      *services.Excecutor
}

func NewInvoiceHandler(invoserv *services.InvoiceService, executor *services.Excecutor) *InvoiceHandler {
	return &InvoiceHandler{InvoiceService: invoserv, Excecutor: executor}
}

func (h *InvoiceHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	// h.InvoiceService.Invoice.DesignInvoice()
	h.Excecutor.Tmpl.ExecuteTemplate(w, "home.page.html", nil)
}
