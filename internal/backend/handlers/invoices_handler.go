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
	if r.Method != http.MethodGet {
		utils.JsonResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// resp, err := h.InvoiceService.List(r.Context(), r)
	// if err != nil {
	// 	log.Println(err)
	// 	utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// h.InvoiceService.Invoice.DesignInvoice()
	h.Excecutor.Tmpl.ExecuteTemplate(w, "home.page.html", nil)
}

func (h *InvoiceHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.JsonResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := h.InvoiceService.ListCustomers(r.Context(), r)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, "customers.page.html", map[string]models.Customers{"Customers": resp}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}

func (h *InvoiceHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.JsonResponse(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := h.InvoiceService.ListProducts(r.Context(), r)
	if err != nil {
		log.Println(err)
		h.Excecutor.Tmpl.ExecuteTemplate(w, "error.page.html", err)
		return
	}

	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, "products.page.html", map[string]models.Products{"Products": resp}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}
