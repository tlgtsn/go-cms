package category

import (
	"cms-project/pkg/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCategoriesHandler handles retrieving all categories
// @Summary Get all categories
// @Description Retrieve all categories
// @Tags Category
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /categories [get]
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := GetAllCategories()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to retrieve categories", nil)
		return
	}
	response.JSON(w, http.StatusOK, true, "Categories retrieved successfully", categories)
}

// CreateCategoryHandler handles creating a new category
// @Summary Create a new category
// @Description Add a new category to the database
// @Tags Category
// @Accept json
// @Produce json
// @Param category body category.CreateCategoryRequest true "Category data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /categories [post]
func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}
	category := Category{
		CreateCategoryRequest: req,
	}
	if err := CreateCategory(category); err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to create category", nil)
		return
	}
	response.JSON(w, http.StatusCreated, true, "Category created successfully", nil)
}

// GetCategoryByIDHandler handles retrieving a single category by ID
// @Summary Get a category by ID
// @Description Retrieve a specific category using its ID
// @Tags Category
// @Param id path int true "Category ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /categories/{id} [get]
func GetCategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid category ID", nil)
		return
	}
	category, err := GetCategoryByID(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, false, "Category not found", nil)
		return
	}
	response.JSON(w, http.StatusOK, true, "Category retrieved successfully", category)
}

// DeleteCategoryHandler handles deleting a category
// @Summary Delete a category
// @Description Remove a category from the database
// @Tags Category
// @Param id path int true "Category ID"
// @Success 204 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /categories/{id} [delete]
func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid category ID", nil)
		return
	}
	if err := DeleteCategory(id); err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to delete category", nil)
		return
	}
	response.JSON(w, http.StatusNoContent, true, "Category deleted successfully", nil)
}
