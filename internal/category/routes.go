package category

import "github.com/gorilla/mux"

// RegisterCategoryRoutes registers all category-related routes
func RegisterCategoryRoutes(r *mux.Router) {
	r.HandleFunc("", GetCategoriesHandler).Methods("GET")                 // List categories
	r.HandleFunc("", CreateCategoryHandler).Methods("POST")               // Create a category
	r.HandleFunc("/{id:[0-9]+}", GetCategoryByIDHandler).Methods("GET")   // Get category by ID
	r.HandleFunc("/{id:[0-9]+}", DeleteCategoryHandler).Methods("DELETE") // Delete a category
}
