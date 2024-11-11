package handlers

import (
	"app-example/internal/dto"
	"app-example/internal/entity"
	"app-example/internal/infra/database"
	pkg_entity "app-example/pkg/entity"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	productDB *database.Product
}

func NewProductHandler(productDB *database.Product) *ProductHandler {
	return &ProductHandler{productDB}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param input body dto.CreateProductInput true "Product data"
// @Success 201 {object} entity.Product
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Erro ao ler o corpo da requisição")
		return
	}
	product, err := entity.NewProduct(productDto.Name, productDto.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	err = h.productDB.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) ListAllProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 8)
	if err != nil {
		page = 0
	}
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 8)
	if err != nil {
		limit = 10
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.productDB.FindAll(int(page), int(limit), sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	if len(products) == 0 {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Nenhum produto encontrado")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := pkg_entity.StringToID(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ID inválido")
		return
	}

	product, err := h.productDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Produto não encontrado")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := pkg_entity.StringToID(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ID inválido")
		return
	}

	var productDto dto.UpdateProductInput
	err = json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Erro ao ler o corpo da requisição")
		return
	}

	product, err := entity.NewProduct(productDto.Name, productDto.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	product.ID = id
	err = h.productDB.Update(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := pkg_entity.StringToID(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ID inválido")
		return
	}

	err = h.productDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
