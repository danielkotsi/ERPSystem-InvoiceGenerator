package invoice

import (
	"invoice_manager/internal/backend/invoice/adapter"
	"invoice_manager/internal/backend/invoice/models"
	"invoice_manager/internal/backend/services"
	"invoice_manager/internal/utils"
	"log"
	"net/http"
)

type InvoiceHandler struct {
	InvoiceService *InvoiceService
	Excecutor      *services.Excecutor
	Adapter        *adapter.InvoiceParser
}

func NewInvoiceHandler(invoserv *InvoiceService, executor *services.Excecutor, adapter *adapter.InvoiceParser) *InvoiceHandler {
	return &InvoiceHandler{InvoiceService: invoserv, Excecutor: executor, Adapter: adapter}
}

func (h *InvoiceHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	h.Excecutor.Tmpl.ExecuteTemplate(w, "home.page.html", nil)
}

func (h *InvoiceHandler) GetMakeInvoicePage(w http.ResponseWriter, r *http.Request) {
	invoiceType := h.Adapter.GetInvoiceTypeFromParameter(r)
	invoiceinfo, err := h.InvoiceService.GetInvoiceInfo(r.Context(), invoiceType)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err, 500)
	}

	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, string(invoiceType.HTMLTemplate()), map[string]models.InvoiceHTMLinfo{"Info": invoiceinfo}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	invo, err := h.Adapter.ParseInvoiceFromRequest(r)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err, 500)
	}
	pdf, err := h.InvoiceService.CreateInvoice(r.Context(), invo)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err, 500)
	}

	utils.PDFResponse(w, pdf, 200)
}
