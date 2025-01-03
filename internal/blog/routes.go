package blog

import (
	"github.com/gorilla/mux"
)

// RegisterBlogRoutes registers all blog routes
func RegisterBlogRoutes(r *mux.Router) {
	r.HandleFunc("", GetBlogsHandler).Methods("GET")
	r.HandleFunc("", CreateBlogHandler).Methods("POST")
	r.HandleFunc("/{id:[a-fA-F0-9-]+}", GetBlogByIDHandler).Methods("GET")
	r.HandleFunc("/{id:[a-fA-F0-9-]+}", UpdateBlogHandler).Methods("PUT")
	r.HandleFunc("/{id:[a-fA-F0-9-]+}", DeleteBlogHandler).Methods("DELETE")
	r.HandleFunc("/search", SearchBlogsHandler).Methods("GET")
	r.HandleFunc("/{id:[a-fA-F0-9-]+}/categories", AddCategoryToBlogHandler).Methods("POST")
	r.HandleFunc("/{id:[a-fA-F0-9-]+}/categories/{category_id:[a-fA-F0-9-]+}", RemoveCategoryFromBlogHandler).Methods("DELETE") // Remove category from blog

}
