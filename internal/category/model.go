package category

import (
	"time"

	"github.com/google/uuid"
)

// Category represents a blog category
type CreateCategoryRequest struct {
	Name        string  `json:"name" example:"Technology"`
	Description *string `json:"description,omitempty" example:"All about technology"`
}

// Category represents a blog category
type Category struct {
	ID                    uuid.UUID        `db:"id" json:"id"`
	CreateCategoryRequest `json:",inline"` // Embed CreateBlogRequest
	CreatedAt             time.Time        `db:"created_at" json:"created_at"`
}
