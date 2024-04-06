package handlers

import (
	"fmt"
	"htmx-demo/components"
	"htmx-demo/db"
	"htmx-demo/models"
	"net/http"
	"strconv"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	var person models.Person
	db := db.Get()

	err := db.First(&person, idString).Error

	if (err != nil) {
		fmt.Println("Couldn't find a person with that ID in the database.")
		return
	}

	components.ClickToEdit(person).Render(r.Context(), w);
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	pageNumberString := r.URL.Query().Get("page")
	pageNumberInt, err := strconv.Atoi(pageNumberString)

	if (err != nil) {
		fmt.Println("Failed to convert page number to an integer")
		return
	}

	nextPageNumberInt := pageNumberInt + 1
	nextPageNumberString := strconv.Itoa(nextPageNumberInt);

	var persons []models.Person
	db := db.Get()
	db.Offset((pageNumberInt - 1) * 10).Limit(10).Find(&persons)

	if (len(persons) == 0) {
		components.NoMoreResults().Render(r.Context(), w);
		return
	}

	components.ClickToLoadMoreResults(persons, nextPageNumberString).Render(r.Context(), w);
}

func PutPerson(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	var person models.Person
	db := db.Get()

	err := db.First(&person, idString).Error

	if (err != nil) {
		fmt.Println("Couldn't find a person with that ID in the database.")
		return
	}

	person.FirstName = r.FormValue("firstName")
	person.LastName = r.FormValue("lastName")
	person.Suburb = r.FormValue("suburb")

	db.Save(&person)
	components.ClickToEdit(person).Render(r.Context(), w);
}

func PutPersonRow(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	var person models.Person
	db := db.Get()

	err := db.First(&person, idString).Error

	if (err != nil) {
		fmt.Println("Couldn't find a person with that ID in the database.")
		return
	}

	person.FirstName = r.FormValue("firstName")
	person.LastName = r.FormValue("lastName")
	person.Suburb = r.FormValue("suburb")

	db.Save(&person)
	components.DeleteEditRow(person).Render(r.Context(), w);
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	db := db.Get()

	db.Delete(&models.Person{}, idString)
}

func GetPersonEdit(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	var person models.Person
	db := db.Get()

	err := db.First(&person, idString).Error

	if (err != nil) {
		fmt.Println("Couldn't find a person with that ID in the database.")
		return
	}

	components.ClickToEditForm(person).Render(r.Context(), w);
}

func GetPersonEditRow(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	var person models.Person
	db := db.Get()

	err := db.First(&person, idString).Error

	if (err != nil) {
		fmt.Println("Couldn't find a person with that ID in the database.")
		return
	}

	components.EditRowForm(person).Render(r.Context(), w);
}

func GetPersonRow(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	var person models.Person
	db := db.Get()

	err := db.First(&person, idString).Error

	if (err != nil) {
		fmt.Println("Couldn't find a person with that ID in the database.")
		return
	}

	components.DeleteEditRow(person).Render(r.Context(), w);
}