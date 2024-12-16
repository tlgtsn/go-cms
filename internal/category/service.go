package category

import (
	"cms-project/internal/database"
	"log"
)

// GetAllCategories retrieves all categories from the database
func GetAllCategories() ([]Category, error) {
	var categories []Category
	query := "SELECT * FROM categories ORDER BY created_at DESC"
	err := database.DB.Select(&categories, query)
	if err != nil {
		log.Printf("Error retrieving categories: %v", err)
		return nil, err
	}
	return categories, nil
}

// CreateCategory inserts a new category into the database
func CreateCategory(category Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, category.Name, category.Description)
	if err != nil {
		log.Printf("Error creating category: %v", err)
		return err
	}
	return nil
}

// GetCategoryByID retrieves a single category by ID
func GetCategoryByID(id int) (*Category, error) {
	var category Category
	query := "SELECT * FROM categories WHERE id = $1"
	err := database.DB.Get(&category, query, id)
	if err != nil {
		log.Printf("Error retrieving category by ID: %v", err)
		return nil, err
	}
	return &category, nil
}

// DeleteCategory deletes a category by ID
func DeleteCategory(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting category: %v", err)
		return err
	}
	return nil
}
