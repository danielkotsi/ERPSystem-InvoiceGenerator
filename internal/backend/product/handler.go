package product

import (
	"-invoice_manager/internal/backend/product/models"
	"-invoice_manager/internal/backend/services"
	"-invoice_manager/internal/utils"
	"log"
	"net/http"
)

type ProductsHandler struct {
	ProductsService *ProductsService
	Excecutor       *services.Excecutor
}

func NewProductsHandler(invoserv *ProductsService, executor *services.Excecutor) *ProductsHandler {
	return &ProductsHandler{ProductsService: invoserv, Excecutor: executor}
}

func (h *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	resp, err := h.ProductsService.ListProducts(r.Context(), search)
	if err != nil {
		log.Println(err)
		h.Excecutor.Tmpl.ExecuteTemplate(w, "error.page.html", err)
		return
	}

	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, "products.page.html", map[string][]models.Product{"Products": resp}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}

func (h *ProductsHandler) GetProductSuggestions(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	resp, err := h.ProductsService.ListProducts(r.Context(), search)
	if err != nil {
		log.Println(err)
		h.Excecutor.Tmpl.ExecuteTemplate(w, "error.page.html", err)
		return
	}

	utils.JsonResponse(w, resp, 200)
}

func (h *ProductsHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := utils.ParseFormData(r, &product); err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err := h.ProductsService.CreateProduct(r.Context(), product)
	if err != nil {
		log.Println(err)
		utils.JsonResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseForClient(w, true, "product created successfully", 200)
}
