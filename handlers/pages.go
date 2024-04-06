package handlers

import (
	"htmx-demo/components"
	"htmx-demo/db"
	"htmx-demo/models"
	"net/http"
)

func PageIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		components.Layout(components.NotFound()).Render(r.Context(), w);
		return
	}
	components.Layout(components.Index()).Render(r.Context(), w);
}

func PageClickToEdit(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	db := db.Get()
	db.First(&person)
	components.Layout(components.ClickToEdit(person)).Render(r.Context(), w);
}

func PageClickToLoad(w http.ResponseWriter, r *http.Request) {
	var persons []models.Person
	db := db.Get()
	db.Limit(10).Find(&persons)
	components.Layout(components.ClickToLoad(persons)).Render(r.Context(), w);
}

func PageDeleteEditRow(w http.ResponseWriter, r *http.Request) {
	var persons []models.Person
	db := db.Get()
	db.Find(&persons)
	components.Layout(components.DeleteEditRows(persons)).Render(r.Context(), w);
}