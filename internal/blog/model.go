package blog

import "time"

// CreateBlogRequest represents the required fields for creating a blog
type CreateBlogRequest struct {
	Title      string  `json:"title" example:"My First Blog"`
	Content    string  `json:"content" example:"This is the content of the blog."`
	Status     string  `db:"status" json:"status"`
	CoverImage *string `db:"cover_image" json:"cover_image,omitempty"`
}

// Blog represents a blog post
type Blog struct {
	ID                int              `db:"id" json:"id"`
	CreateBlogRequest `json:",inline"` // Embed CreateBlogRequest
	AuthorID          int              `db:"author_id" json:"author_id"`
	CreatedAt         time.Time        `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time        `db:"updated_at" json:"updated_at"`
}

// BlogCategory represents the relationship between blogs and categories
type BlogCategory struct {
	BlogID     int `db:"blog_id"`
	CategoryID int `db:"category_id"`
}
