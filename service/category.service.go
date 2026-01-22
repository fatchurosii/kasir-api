package service

import (
	"errors"
	"kasir-api/data"
	"kasir-api/entity"
)

var ErrCategoryNotFound = errors.New("Category not found")

func GetCategoryByID(id int) (entity.Category, error) {
	for _, category := range data.Categories {
		if category.ID == id {
			return category, nil
		}
	}

	return entity.Category{}, ErrCategoryNotFound
}

func GetAllCategory() []entity.Category {
	return data.Categories
}

func CreateCategory(category entity.Category) entity.Category {
	category.ID = len(data.Categories) + 1
	data.Categories = append(data.Categories, category)
	return category
}

func UpdateCategory(id int, category entity.Category) (entity.Category, error) {
	for i, p := range data.Categories {
		if p.ID == id {
			category.ID = id
			data.Categories[i] = category
			return category, nil
		}
	}
	return entity.Category{}, ErrCategoryNotFound
}

func DeleteCategory(id int) (entity.Category, error) {
	for i, category := range data.Categories {
		if category.ID == id {
			data.Categories = append(
				data.Categories[:i],
				data.Categories[i+1:]...,
			)
			return category, nil
		}
	}
	return entity.Category{}, ErrCategoryNotFound
}
