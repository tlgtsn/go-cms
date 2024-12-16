package category

import (
	"time"
)

// Category represents a blog category
type CreateCategoryRequest struct {
	Name        string  `db:"name" json:"name"`
	Description *string `db:"description" json:"description,omitempty"`
}

// Category represents a blog category
type Category struct {
	ID                    int              `db:"id" json:"id"`
	CreateCategoryRequest `json:",inline"` // Embed CreateBlogRequest
	CreatedAt             time.Time        `db:"created_at" json:"created_at"`
}
