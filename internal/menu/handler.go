package menu

import (
	"cms-project/pkg/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetMenusHandler handles retrieving all menus
// @Summary Get all menus
// @Description Retrieve all menus, optionally filter by parent_id
// @Tags Menu
// @Param parent_id query int false "Parent menu ID"
// @Success 200 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /menus [get]
func GetMenusHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	menus, err := GetMenus(page, limit)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to fetch menus", nil)

		return
	}
	response.JSON(w, http.StatusOK, true, "Menus retrieved successfully", menus)
}

// CreateMenuHandler handles creating a new menu
// @Summary Create a new menu
// @Description Add a new menu to the database
// @Tags Menu
// @Param menu body menu.CreateMenuRequest true "Menu data"
// @Success 201 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /menus [post]
func CreateMenuHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateMenuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}
	menu := Menu{
		CreateMenuRequest: req,
	}
	if err := CreateMenu(menu); err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to create menu", nil)
		return
	}
	response.JSON(w, http.StatusCreated, true, "Menu created successfully", nil)
}

// GetMenuByIDHandler handles retrieving a single menu by ID
// @Summary Get a menu by ID
// @Description Retrieve a specific menu using its ID
// @Tags Menu
// @Param id path int true "Menu ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Router /menus/{id} [get]
func GetMenuByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid menu ID", nil)
		return
	}
	menu, err := GetMenuByID(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, false, "Menu not found", nil)

		return
	}
	response.JSON(w, http.StatusOK, true, "Menu retrieved successfully", menu)
}

// DeleteMenuHandler handles deleting a menu by ID
// @Summary Delete a menu
// @Description Remove a menu from the database
// @Tags Menu
// @Param id path int true "Menu ID"
// @Success 204 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /menus/{id} [delete]
func DeleteMenuHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid menu ID", nil)
		return
	}
	if err := DeleteMenu(id); err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to delete menu", nil)
		return
	}
	response.JSON(w, http.StatusNoContent, true, "Menu delete successfully", nil)
}

// @Summary Update a menu
// @Description Update a menu's name or parent_id using its ID
// @Tags Menu
// @Param id path int true "Menu ID"
// @Param menu body menu.CreateMenuRequest true "Menu data to update"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 404 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /menus/{id} [put]
func UpdateMenuHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid menu Id", nil)
		return
	}

	var req CreateMenuRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.JSON(w, http.StatusBadRequest, false, "Invalid input", nil)
		return
	}

	menu := Menu{
		ID:                id,
		CreateMenuRequest: req,
	}

	if err := UpdateMenu(menu); err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to update menu", nil)
		return
	}
	response.JSON(w, http.StatusOK, true, "Menu updated successfully", nil)
}

// FilterMenusHandler handles filtering menus by parent_id
// @Summary Filter menus
// @Description Retrieve menus filtered by parent_id
// @Tags Menu
// @Param parent_id query int false "Parent menu ID"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /menus/filter [get]
func FilterMenusHandler(w http.ResponseWriter, r *http.Request) {
	// Query parameter
	parentIDStr := r.URL.Query().Get("parent_id")
	var parentID *int
	if parentIDStr != "" {
		id, err := strconv.Atoi(parentIDStr)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, false, "Invalid parent_id", nil)
			return
		}
		parentID = &id
	}

	// Call service
	menus, err := FilterMenus(parentID)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, false, "Failed to filter menus", nil)
		return
	}

	// Success response
	response.JSON(w, http.StatusOK, true, "Menus retrieved successfully", menus)
}
