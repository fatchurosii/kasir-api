package repository

import (
	"database/sql"
	"errors"
	"kasir-api/entity"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repo *CategoryRepository) GetAllCategories() ([]entity.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]entity.Category, 0)
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (repo *CategoryRepository) GetCategoryByID(id int) (*entity.Category, error) {

	query := `SELECT id, name, description FROM categories WHERE id = $1`

	var category entity.Category
	err := repo.db.QueryRow(query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("category not found")
	}
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (repo *CategoryRepository) CreateCategory(category *entity.Category) error {
	query := `INSERT INTO categories (name, description) VALUES ($1,$2)`
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)

	if err != nil {
		return err
	}
	return err
}

func (repo *CategoryRepository) UpdateCategory(category *entity.Category) error {
	query := `UPDATE categories SET name = $1, price = $2 , description = $3 WHERE id = $4`

	result, err := repo.db.Exec(query, category.Name, category.Description, category.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}

func (repo *CategoryRepository) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id = $1`
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("category not found")
	}
	return nil
}
