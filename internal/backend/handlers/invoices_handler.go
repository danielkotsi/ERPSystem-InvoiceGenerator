package handlers

import (
	"-invoice_manager/internal/backend/services"
	"-invoice_manager/internal/utils"
	"net/http"
)

type InvoiceHandler struct {
	InvoiceService *services.InvoiceService
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

	h.InvoiceService.Invoice.DesignInvoice()
	h.InvoiceService.Tmpl.ExecuteTemplate(w, "home.page.html", nil)
}
