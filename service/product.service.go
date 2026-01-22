package service

import (
	"errors"
	"kasir-api/data"
	"kasir-api/entity"
)

var ProductNotFound = errors.New("product not found")

func GetProductByID(id int) (entity.Product, error) {
	for _, product := range data.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return entity.Product{}, ProductNotFound
}

func GetAllProduct() []entity.Product {
	return data.Products
}

func CreateProduct(product entity.Product) entity.Product {
	product.ID = len(data.Products) + 1
	data.Products = append(data.Products, product)
	return product
}

func UpdateProduct(id int, product entity.Product) (entity.Product, error) {
	for i, p := range data.Products {
		if p.ID == id {
			product.ID = id
			data.Products[i] = product
			return product, nil
		}
	}
	return entity.Product{}, ProductNotFound
}

func DeleteProduct(id int) (entity.Product, error) {
	for i, product := range data.Products {
		if product.ID == id {
			data.Products = append(
				data.Products[:i],
				data.Products[i+1:]...,
			)
			return product, nil
		}
	}
	return entity.Product{}, ProductNotFound
}
