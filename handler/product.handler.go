package handler

import (
	"kasir-api/entity"
	"kasir-api/helper"
	"kasir-api/service"
	"net/http"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/product/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	Success(w, http.StatusOK, "Product found", product)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProductsService()
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "All products", products)
}

func (h *ProductHandler) StoreProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	if err := helper.DecodeJSON(r, &product); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	result := h.service.CreateProduct(&product)
	Success(w, http.StatusCreated, "Product created", result)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/product/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	var product entity.Product
	if err := helper.DecodeJSON(r, &product); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}
	product.ID = id
	err = h.service.UpdateProduct(&product)
	if err != nil {
		Error(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	Success(w, http.StatusOK, "Product updated", product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/product/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	err = h.service.ProductRepo.DeleteProduct(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	Success(w, http.StatusOK, "Product deleted", nil)
}
