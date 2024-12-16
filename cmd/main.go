package main

import (
	_ "cms-project/internal/blog" // Swagger için gerekli
	"cms-project/internal/database"
	"cms-project/internal/routes"
	"log"
	"net/http"

	_ "cms-project/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title CMS Project API
// @version 1.0
// @description This is a simple CMS API with blog and menu features.
// @host localhost:8080
// @BasePath /

// @contact.name Your Name
// @contact.url https://your-website.com
// @contact.email your-email@example.com
func main() {
	// Initialize database
	database.InitDB()

	r := routes.InitializeRoutes()
	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start the server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
