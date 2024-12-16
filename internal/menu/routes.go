package menu

import (
	"github.com/gorilla/mux"
)

// RegisterMenuRoutes registers all menu routes
func RegisterMenuRoutes(r *mux.Router) {
	r.HandleFunc("", GetMenusHandler).Methods("GET")
	r.HandleFunc("", CreateMenuHandler).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", GetMenuByIDHandler).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", UpdateMenuHandler).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}", DeleteMenuHandler).Methods("DELETE")
	r.HandleFunc("/filter", FilterMenusHandler).Methods("GET")
}
