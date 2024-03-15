package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gergosabian/go-message/pkg/router"
)

func main() {
	// Setup router
	r := router.SetupRouter()

	// Start server
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
