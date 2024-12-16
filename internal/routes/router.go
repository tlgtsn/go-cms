package routes

import (
	"cms-project/internal/blog"
	"cms-project/internal/category"
	"cms-project/internal/menu"

	"github.com/gorilla/mux"
)

// InitializeRoutes initializes all application routes
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	// Blog routes
	blogRouter := r.PathPrefix("/blogs").Subrouter()
	blog.RegisterBlogRoutes(blogRouter)

	// Menu routes
	menuRouter := r.PathPrefix("/menus").Subrouter()
	menu.RegisterMenuRoutes(menuRouter)

	// Category routes
	categoryRouter := r.PathPrefix("/categories").Subrouter()
	category.RegisterCategoryRoutes(categoryRouter)

	return r
}
