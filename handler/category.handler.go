package handler

import (
	"kasir-api/entity"
	"kasir-api/helper"
	"kasir-api/service"
	"net/http"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/category/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	category, err := h.service.GetCategoryByID(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Category not found", nil)
		return
	}

	Success(w, http.StatusOK, "Category found", category)
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAllCategory()
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	Success(w, http.StatusOK, "All categories", categories)
}

func (h *CategoryHandler) StoreCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	if err := helper.DecodeJSON(r, &category); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	result := h.service.CreateCategory(&category)
	Success(w, http.StatusCreated, "Category created", result)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/category/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	var category entity.Category
	if err := helper.DecodeJSON(r, &category); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}
	category.ID = id
	err = h.service.UpdateCategory(&category)
	if err != nil {
		Error(w, http.StatusNotFound, "Category not found", nil)
		return
	}

	Success(w, http.StatusOK, "Category updated", category)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/category/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	err = h.service.CategoryRepo.DeleteCategory(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Category not found", nil)
		return
	}

	Success(w, http.StatusOK, "Category deleted", nil)
}
