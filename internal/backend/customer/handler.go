package customer

import (
	"-invoice_manager/internal/backend/customer/models"
	"-invoice_manager/internal/backend/invoice/payload"
	"-invoice_manager/internal/backend/services"
	"-invoice_manager/internal/utils"
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

// excecutes html for customers page, uses the customer repo to get all the info for all the customers
// Need to add a limit for the cusotmers returned and maybe choose less info to be returned
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

// That's done doesnt need to change
// ececute the html for the customer by id uses the customer repo to get all the info for one customer based on the codeNumber
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

// responds with Json, returns all the customers with all the info for the suggestions to render and also fill the fields when they are chosen
// needs to be refactored so that it returns only the name.
func (h *CustomersHandler) GetCustomerSuggestions(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	resp, err := h.CustomersService.ListCustomerSuggestions(r.Context(), search)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, resp, 200)
}

// responds with Json, returns all the info for all the branch Companies of a certain company
// needs to be refactored so that it returns only the name and the codeNumber
func (h *CustomersHandler) GetBranchCompaniesSuggestions(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	company := r.URL.Query().Get("company")
	resp, err := h.CustomersService.ListBranchCompanySuggestions(r.Context(), search, company)
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
