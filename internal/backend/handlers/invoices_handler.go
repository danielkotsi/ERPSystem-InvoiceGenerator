package handlers

import (
	"-invoice_manager/internal/backend/services"
	"log"
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

func (h *InvoiceHandler) GetMakeInvoicePage(w http.ResponseWriter, r *http.Request) {
	// h.InvoiceService.Invoice.DesignInvoice()
	h.Excecutor.Tmpl.ExecuteTemplate(w, "create_invoice.page.html", nil)
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	pdf, err := h.InvoiceService.CreateInvoice(r.Context(), r)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=\"document.pdf\"")
	w.Write(pdf)
}
