package menu

import (
	"cms-project/internal/database"
	"log"
)

// GetMenus retrieves all menus from the database
func GetMenus(page, limit int) ([]Menu, error) {
	var menus []Menu
	offset := (page - 1) * limit
	query := "SELECT * FROM menus ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	err := database.DB.Select(&menus, query, limit, offset)
	if err != nil {
		log.Printf("Error fetching menus: %v", err)
		return nil, err
	}
	return menus, nil
}

// CreateMenu inserts a new menu into the database
func CreateMenu(menu Menu) error {
	query := "INSERT INTO menus (name, parent_id) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, menu.Name, menu.ParentID)
	if err != nil {
		log.Printf("Error creating menu: %v", err)
		return err
	}
	return nil
}

// GetMenuByID retrieves a single menu by its ID
func GetMenuByID(id int) (*Menu, error) {
	var menu Menu
	query := "SELECT * FROM menus WHERE id = $1"
	err := database.DB.Get(&menu, query, id)
	if err != nil {
		log.Printf("Error fetching menu by ID: %v", err)
		return nil, err
	}
	return &menu, nil
}

// DeleteMenu removes a menu from the database
func DeleteMenu(id int) error {
	query := "DELETE FROM menus WHERE id = $1"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting menu: %v", err)
		return err
	}
	return nil
}

// UpdateMenu updates an existing menu
func UpdateMenu(menu Menu) error {
	query := "UPDATE menus SET name = $1, parent_id ? $2 WHERE id = $3"
	_, err := database.DB.Exec(query, menu.Name, menu.ParentID, menu.ID)
	if err != nil {
		log.Printf("Error updating menu: %v", err)
		return err
	}
	return nil
}

// FilterMenus filters menus by parent_id
func FilterMenus(parentID *int) ([]Menu, error) {
	var menus []Menu

	query := "SELECT * FROM menus WHERE parent_id = $1 ORDER BY created_at DESC"
	err := database.DB.Select(&menus, query, parentID)
	if err != nil {
		log.Printf("Error filtering menus: %v", err)
		return nil, err
	}
	return menus, nil
}
