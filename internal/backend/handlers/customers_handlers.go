package handlers

import (
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/services"
	"-invoice_manager/internal/utils"
	"log"
	"net/http"
)

type CustomersHandler struct {
	CustomersService *services.CustomersService
	Excecutor        *services.Excecutor
}

func NewCustomersHandler(invoserv *services.CustomersService, executor *services.Excecutor) *CustomersHandler {
	return &CustomersHandler{CustomersService: invoserv, Excecutor: executor}
}

func (h *CustomersHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {

	resp, err := h.CustomersService.ListCustomers(r.Context(), r)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, "customers.page.html", map[string]models.Customers{"Customers": resp}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}

func (h *CustomersHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	err := h.CustomersService.CreateCustomer(r.Context(), r)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseForClient(w, true, "customer created successfully", 200)
}
