package handlers

import (
	"-invoice_manager/internal/backend/models"
	"-invoice_manager/internal/backend/services"
	"log"
	"net/http"
)

type ProductsHandler struct {
	ProductsService *services.ProductsService
	Excecutor       *services.Excecutor
}

func NewProductsHandler(invoserv *services.ProductsService, executor *services.Excecutor) *ProductsHandler {
	return &ProductsHandler{ProductsService: invoserv, Excecutor: executor}
}

func (h *ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	resp, err := h.ProductsService.ListProducts(r.Context(), r)
	if err != nil {
		log.Println(err)
		h.Excecutor.Tmpl.ExecuteTemplate(w, "error.page.html", err)
		return
	}

	if err := h.Excecutor.Tmpl.ExecuteTemplate(w, "products.page.html", map[string]models.Products{"Products": resp}); err != nil {
		h.Excecutor.ServeErrorwithHTML(w, err, 500)
	}
}
