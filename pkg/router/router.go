package router

import (
	"html/template"
	"net/http"

	"github.com/gergosabian/go-message/pkg/handlers"
)

// SetupRouter sets up the router for the API.
func SetupRouter() http.Handler {
	router := http.NewServeMux()

	// User routes
	router.HandleFunc("GET /users", handlers.GetUsers)
	router.HandleFunc("GET /users/{id}", handlers.GetUser)
	router.HandleFunc("POST /users", handlers.CreateUser)
	router.HandleFunc("PUT /users/{id}", handlers.UpdateUser)
	router.HandleFunc("DELETE /users/{id}", handlers.DeleteUser)

	// Serve static files
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// View resolver
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("pkg/templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return router
}
