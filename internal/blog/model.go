package blog

import (
	"time"

	"github.com/google/uuid"
)

// CreateBlogRequest represents the required fields for creating a blog
type CreateBlogRequest struct {
	Title      string `json:"title" example:"My First Blog"`
	Content    string `json:"content" example:"This is the content of the blog."`
	Status     string `json:"status" example:"draft"` // draft, published
	CoverImage string `json:"cover_image,omitempty" example:"https://example.com/image.jpg"`
	AuthorID   string `json:"author_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
}

// Blog represents a blog post
type Blog struct {
	ID uuid.UUID `db:"id" json:"id"`
	CreateBlogRequest
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// BlogCategory represents the relationship between blogs and categories
type BlogCategory struct {
	BlogID     uuid.UUID `db:"blog_id"`
	CategoryID uuid.UUID `db:"category_id"`
}
