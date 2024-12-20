package blog

import (
	"cms-project/internal/database"
	"fmt"
	"log"

	"github.com/google/uuid"
)

// GetBlogs retrieves all blogs from the database
func GetBlogs(page, limit int) ([]Blog, error) {
	var blogs []Blog

	offset := (page - 1) * limit
	query := "SELECT * FROM blogs ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	err := database.DB.Select(&blogs, query, limit, offset)
	if err != nil {
		log.Printf("Error fetching blogs: %v", err)
		return nil, err
	}
	return blogs, nil
}

// CreateBlog inserts a new blog into the database
func CreateBlog(blog Blog) error {
	query := "INSERT INTO blogs (id, title, content, status, cover_image, author_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())"
	blog.ID = uuid.New()
	_, err := database.DB.Exec(query, blog.ID, blog.Title, blog.Content, blog.Status, blog.CoverImage, blog.AuthorID)
	if err != nil {
		log.Printf("Error creating blog: %v", err)
		return err
	}
	return nil
}

// GetBlogByID retrieves a single blog by its ID
func GetBlogByID(id uuid.UUID) (*Blog, error) {
	var blog Blog
	query := "SELECT * FROM blogs WHERE id = $1"
	err := database.DB.Get(&blog, query, id)
	if err != nil {
		log.Printf("Error fetching blog by ID: %v", err)
		return nil, err
	}
	return &blog, nil
}

// DeleteBlog removes a blog from the database
func DeleteBlog(id uuid.UUID) error {
	query := "DELETE FROM blogs WHERE id = $1"
	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting blog: %v", err)
		return err
	}
	return nil
}

// UpdateBlog updates an existing blog
func UpdateBlog(blog Blog) error {
	query := "UPDATE blogs SET title = $1, content = $2, status = $3, cover_image = $4,  updated_at = NOW() WHERE id = $5"
	_, err := database.DB.Exec(query, blog.Title, blog.Content, blog.CoverImage, blog.Status, blog.ID)
	if err != nil {
		log.Printf("Error updating blog: %v", err)
		return err
	}
	return nil
}

// SearchBlogs searches blogs by title or content
func SearchBlogs(keyword string, page, limit int) ([]Blog, error) {
	var blogs []Blog

	offset := (page - 1) * limit
	query := `
		SELECT * FROM blogs 
		WHERE title ILIKE $1 OR content ILIKE $1
		ORDER BY created_at DESC 
		LIMIT $2 OFFSET $3`
	err := database.DB.Select(&blogs, query, fmt.Sprintf("%%%s%%", keyword), limit, offset)
	if err != nil {
		log.Printf("Error searching blogs: %v", err)
		return nil, err
	}
	return blogs, nil
}

// AddCategoryToBlog adds a category to a blog
func AddCategoryToBlog(blogID, categoryID int) error {
	query := "INSERT INTO blog_categories (blog_id, category_id) VALUES ($1, $2) ON CONFLICT DO NOTHING"
	_, err := database.DB.Exec(query, blogID, categoryID)
	if err != nil {
		log.Printf("Error adding category to blog: %v", err)
		return err
	}
	return nil
}

// RemoveCategoryFromBlog removes a category from a blog
func RemoveCategoryFromBlog(blogID, categoryID uuid.UUID) error {
	query := "DELETE FROM blog_categories WHERE blog_id = $1 AND category_id = $2"
	_, err := database.DB.Exec(query, blogID, categoryID)
	if err != nil {
		log.Printf("Error removing category from blog: %v", err)
		return err
	}
	return nil
}
