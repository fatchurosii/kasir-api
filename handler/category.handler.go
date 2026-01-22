package handler

import (
	"kasir-api/entity"
	"kasir-api/helper"
	"kasir-api/service"
	"net/http"
)

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/category/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	product, err := service.GetCategoryByID(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Category not found", nil)
		return
	}

	Success(w, http.StatusOK, "Category found", product)
}

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	Success(w, http.StatusOK, "All categories", service.GetAllCategory())
}

func StoreCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	if err := helper.DecodeJSON(r, &category); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	result := service.CreateCategory(category)
	Success(w, http.StatusCreated, "Category created", result)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/category/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid product ID", err.Error())
		return
	}

	var category entity.Category
	if err := helper.DecodeJSON(r, &category); err != nil {
		Error(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	result, err := service.UpdateCategory(id, category)
	if err != nil {
		Error(w, http.StatusNotFound, "Product not found", nil)
		return
	}

	Success(w, http.StatusOK, "Product updated", result)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ParseID(r, "/api/category/")
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid category ID", err.Error())
		return
	}

	product, err := service.DeleteCategory(id)
	if err != nil {
		Error(w, http.StatusNotFound, "Category not found", nil)
		return
	}

	Success(w, http.StatusOK, "Category deleted", product)
}
