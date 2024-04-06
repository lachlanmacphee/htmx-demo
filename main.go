package main

import (
	"fmt"
	"htmx-demo/db"
	"htmx-demo/handlers"
	"net/http"
)

func main() {

	db.Connect()
	db.Migrate()
	db.Seed()

	mux := &http.ServeMux{}

	// Public Files
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/favicon.ico")
	})

	// Pages
	mux.HandleFunc("GET /", handlers.PageIndex)
	mux.HandleFunc("GET /demo/click-to-edit", handlers.PageClickToEdit)
	mux.HandleFunc("GET /demo/click-to-load", handlers.PageClickToLoad)
	mux.HandleFunc("GET /demo/delete-edit-row", handlers.PageDeleteEditRow)

	// Person routes
	mux.HandleFunc("GET /person", handlers.GetPersons)
	mux.HandleFunc("GET /person/{id}", handlers.GetPerson)
	mux.HandleFunc("PUT /person/{id}", handlers.PutPerson)
	mux.HandleFunc("DELETE /person/{id}", handlers.DeletePerson)
	mux.HandleFunc("GET /person/{id}/edit", handlers.GetPersonEdit)
	mux.HandleFunc("GET /person/{id}/row", handlers.GetPersonRow)
	mux.HandleFunc("PUT /person/{id}/row", handlers.PutPersonRow)
	mux.HandleFunc("GET /person/{id}/edit-row", handlers.GetPersonEditRow)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}