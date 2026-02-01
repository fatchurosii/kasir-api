package service

import (
	"errors"
	"kasir-api/entity"
	"kasir-api/repository"
)

type CategoryService struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepo: categoryRepo,
	}
}

var ErrCategoryNotFound = errors.New("Category not found")

func (s *CategoryService) GetCategoryByID(id int) (*entity.Category, error) {
	return s.CategoryRepo.GetCategoryByID(id)
}

func (s *CategoryService) GetAllCategory() ([]entity.Category, error) {
	return s.CategoryRepo.GetAllCategories()
}

func (s *CategoryService) CreateCategory(category *entity.Category) error {
	return s.CategoryRepo.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(category *entity.Category) error {
	return s.CategoryRepo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.CategoryRepo.DeleteCategory(id)
}
