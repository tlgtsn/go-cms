package blog

import (
	"cms-project/pkg/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetBlogsHandler handles retrieving all blogs
// @Summary Get all blogs
// @Description Retrieve all blogs with pagination
// @Tags Blog
// @Param page query int false "Page number"
// @Param limit query int false "Number of blogs per page"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /blogs [get]
func GetBlogsHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	blogs, err := GetBlogs(page, limit)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to fetch blogs", nil)
		return
	}
	response.JSON(w, http.StatusOK, true, "Blogs retrieved successfully", blogs)
}

// CreateBlogHandler handles creating a new blog
// @Summary Create anew blog
// @Description Add a new blog with title and content to the database
// @Tags Blog
// @Param blog body blog.CreateBlogRequest  true "Blog data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /blogs [post]
func CreateBlogHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid JSON input", nil)
		return
	}

	blog := Blog{
		CreateBlogRequest: req,
	}

	if err := CreateBlog(blog); err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to create blog", nil)
		return
	}
	response.JSON(w, http.StatusCreated, true, "Blog created successfully", nil)
}

// GetBlogByIDHandler handles retrieving a single blog by ID
// @Summary Get a blog by ID
// @Description Retrieve a specific blog using its ID
// @Tags Blog
// @Param id path int true "Blog ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /blogs/{id} [get]
func GetBlogByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid blog ID", nil)
		return
	}
	blog, err := GetBlogByID(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, false, "Blog not found", nil)
		return
	}
	response.JSON(w, http.StatusOK, true, "Blog retrieved successfully", blog)
}

// DeleteBlogHandler handles deleting a blog by ID
// @Summary Delete a blog
// @Description Remove a blog from the database
// @Tags Blog
// @Param id path int true "Blog ID"
// @Success 204 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /blogs/{id} [delete]
func DeleteBlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid blog ID", nil)
		return
	}
	if err := DeleteBlog(id); err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to delete blog", nil)
		return
	}
	response.JSON(w, http.StatusCreated, true, "Blog delete successfully", nil)
}

// @Summary Update a blog
// @Description Update a blog's title and content using its ID
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Param blog body blog.CreateBlogRequest  true "Blog data to update"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /blogs/{id} [put]
func UpdateBlogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid blog ID", nil)
		return
	}

	var req CreateBlogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}

	blog := Blog{
		ID:                id,
		CreateBlogRequest: req,
	}

	if err := UpdateBlog(blog); err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to update blog", nil)
		return
	}

	response.JSON(w, http.StatusOK, true, "Blog updated successfully", nil)
}

// SearchBlogsHandler handles searching blogs
// @Summary Search blogs
// @Description Search blogs by title or content using a keyword
// @Tags Blog
// @Param keyword query string true "Keyword to search for"
// @Param page query int false "Page number"
// @Param limit query int false "Number of blogs per page"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /blogs/search [get]
func SearchBlogsHandler(w http.ResponseWriter, r *http.Request) {
	// Query parameters
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		response.JSON(w, http.StatusBadRequest, false, "Keyword is required", nil)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Call service
	blogs, err := SearchBlogs(keyword, page, limit)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to search blogs", nil)
		return
	}

	// Success response
	response.JSON(w, http.StatusOK, true, "Blogs retrieved successfully", blogs)
}
