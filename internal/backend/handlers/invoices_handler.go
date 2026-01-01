package handlers

import (
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/services"
	"-invoice_manager/internal/utils"
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

	invoiceinfo, invoicehtml, err := h.InvoiceService.GetInvoiceInfo(r.Context(), r)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err, 500)
	}
	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, invoicehtml, map[string]models.InvoiceHTMLinfo{"Info": invoiceinfo}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	pdf, err := h.InvoiceService.CreateInvoice(r.Context(), r)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err, 500)
	}

	// if err := h.Excecutor.Tmpl.ExecuteTemplate(w, "invoice.page.html", map[string]models.Invoice{"Customers": pdf}); err != nil {
	// 	h.Excecutor.ServeErrorwithHTML(w, err, 500)
	// }
	utils.PDFResponse(w, pdf, 200)
	// utils.XMLResponse(w, pdf, 200)
}
