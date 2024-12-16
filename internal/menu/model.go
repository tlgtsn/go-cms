package menu

import "time"

// CreateMenuRequest represents the required fields for creating a menu
type CreateMenuRequest struct {
	Name     string `db:"name" json:"name" example:"Main Menu"`
	ParentID *int   `db:"parent_id,omitempty" json:"parent_id,omitempty" example:"1"`
}

// Menu represents a menu item
type Menu struct {
	ID                int              `db:"id" json:"id"`
	CreateMenuRequest `json:",inline"` // Embed CreateMenuRequest
	CreatedAt         time.Time        `db:"created_at" json:"created_at"`
}
