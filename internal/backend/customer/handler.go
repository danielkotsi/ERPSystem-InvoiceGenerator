package customer

import (
	"invoice_manager/internal/backend/customer/models"
	"invoice_manager/internal/backend/invoice/payload"
	"invoice_manager/internal/backend/services"
	"invoice_manager/internal/utils"
	"log"
	"net/http"
	"strings"
)

type CustomersHandler struct {
	CustomersService *CustomersService
	Excecutor        *services.Excecutor
}

func NewCustomersHandler(invoserv *CustomersService, executor *services.Excecutor) *CustomersHandler {
	return &CustomersHandler{CustomersService: invoserv, Excecutor: executor}
}

func (h *CustomersHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	resp, err := h.CustomersService.ListCustomers(r.Context(), search)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, "customers.page.html", map[string][]payload.Company{"Customers": resp}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}

func (h *CustomersHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	codeNumber := strings.TrimPrefix(r.URL.String(), "/customers/byid/")
	resp, err := h.CustomersService.GetCustomerById(r.Context(), codeNumber)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, "customerbyid.page.html", map[string]models.CustomerById{"Resp": resp}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}

func (h *CustomersHandler) GetCustomerSuggestions(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	resp, err := h.CustomersService.ListCustomers(r.Context(), search)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, resp, 200)
}

func (h *CustomersHandler) GetBranchCompaniesSuggestions(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	company := r.URL.Query().Get("company")
	resp, err := h.CustomersService.ListBranchCompanies(r.Context(), search, company)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, resp, 200)
}

func (h *CustomersHandler) CreateBranchCompany(w http.ResponseWriter, r *http.Request) {
	var branch payload.BranchCompany
	if err := utils.ParseFormData(r, &branch); err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := h.CustomersService.CreateBranchCompany(r.Context(), branch); err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseForClient(w, true, "branch company created successfully", 200)
}

func (h *CustomersHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer payload.Company
	if err := utils.ParseFormData(r, &customer); err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := h.CustomersService.CreateCustomer(r.Context(), customer)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseForClient(w, true, "customer created successfully", 200)
}
