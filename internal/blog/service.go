package blog

import (
	"cms-project/internal/database"
	"fmt"
	"log"
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
	query := "INSERT INTO blogs (title, content) VALUES ($1, $2)"
	_, err := database.DB.Exec(query, blog.Title, blog.Content)
	if err != nil {
		log.Printf("Error creating blog: %v", err)
		return err
	}
	return nil
}

// GetBlogByID retrieves a single blog by its ID
func GetBlogByID(id int) (*Blog, error) {
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
func DeleteBlog(id int) error {
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
	query := "UPDATE blogs SET title = $1, content = $2 WHERE id = $3"
	_, err := database.DB.Exec(query, blog.Title, blog.Content, blog.ID)
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
