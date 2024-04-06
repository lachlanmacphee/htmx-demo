package db

import (
	"htmx-demo/models"
	"log"
)

func Seed() {
    db := Get() // Get the database connection

    // Create seed data
    people := []models.Person{
        {FirstName: "John", LastName: "Doe", StreetNumber: "123", StreetName: "Main St", Suburb: "Downtown", City: "Cityville", State: "CA", Country: "USA"},
        {FirstName: "Jane", LastName: "Smith", StreetNumber: "456", StreetName: "Oak Ave", Suburb: "Oakville", City: "Townsville", State: "TX", Country: "USA"},
        {FirstName: "Bob", LastName: "Johnson", StreetNumber: "789", StreetName: "Elm Rd", Suburb: "Elmtown", City: "Villageville", State: "NY", Country: "USA"},
        {FirstName: "Alice", LastName: "Williams", StreetNumber: "101", StreetName: "Maple Dr", Suburb: "Maplewood", City: "Oaktown", State: "CA", Country: "USA"},
        {FirstName: "Michael", LastName: "Brown", StreetNumber: "222", StreetName: "Pine Ln", Suburb: "Pineville", City: "Elmtown", State: "TX", Country: "USA"},
        {FirstName: "Emily", LastName: "Davis", StreetNumber: "333", StreetName: "Cedar Rd", Suburb: "Cedarville", City: "Oakville", State: "NY", Country: "USA"},
        {FirstName: "David", LastName: "Miller", StreetNumber: "444", StreetName: "Birch Ave", Suburb: "Birchwood", City: "Mapleton", State: "CA", Country: "USA"},
        {FirstName: "Sarah", LastName: "Wilson", StreetNumber: "555", StreetName: "Willow St", Suburb: "Willowbrook", City: "Pineton", State: "TX", Country: "USA"},
        {FirstName: "Daniel", LastName: "Moore", StreetNumber: "666", StreetName: "Ash Ln", Suburb: "Ashville", City: "Cedarton", State: "NY", Country: "USA"},
        {FirstName: "Jessica", LastName: "Taylor", StreetNumber: "777", StreetName: "Oak Blvd", Suburb: "Oakridge", City: "Birchtown", State: "CA", Country: "USA"},
        {FirstName: "Matthew", LastName: "Anderson", StreetNumber: "888", StreetName: "Maple Rd", Suburb: "Maplecrest", City: "Willowville", State: "TX", Country: "USA"},
        {FirstName: "Ashley", LastName: "Thomas", StreetNumber: "999", StreetName: "Pine Ave", Suburb: "Pinecrest", City: "Ashton", State: "NY", Country: "USA"},
        {FirstName: "Joshua", LastName: "Jackson", StreetNumber: "111", StreetName: "Cedar St", Suburb: "Cedarbrook", City: "Oakridge", State: "CA", Country: "USA"},
        {FirstName: "Amanda", LastName: "White", StreetNumber: "222", StreetName: "Birch Ln", Suburb: "Birchfield", City: "Maplecrest", State: "TX", Country: "USA"},
        {FirstName: "Christopher", LastName: "Harris", StreetNumber: "333", StreetName: "Willow Rd", Suburb: "Willowfield", City: "Pinecrest", State: "NY", Country: "USA"},
        {FirstName: "Stephanie", LastName: "Martin", StreetNumber: "444", StreetName: "Ash Ave", Suburb: "Ashbrook", City: "Cedarbrook", State: "CA", Country: "USA"},
        {FirstName: "Jacob", LastName: "Thompson", StreetNumber: "555", StreetName: "Oak St", Suburb: "Oakdale", City: "Birchfield", State: "TX", Country: "USA"},
        {FirstName: "Brittany", LastName: "Garcia", StreetNumber: "666", StreetName: "Maple Blvd", Suburb: "Maplegrove", City: "Willowfield", State: "NY", Country: "USA"},
        {FirstName: "Andrew", LastName: "Martinez", StreetNumber: "777", StreetName: "Pine Rd", Suburb: "Pinecrest", City: "Ashbrook", State: "CA", Country: "USA"},
        {FirstName: "Samantha", LastName: "Robinson", StreetNumber: "888", StreetName: "Cedar Ave", Suburb: "Cedarville", City: "Oakdale", State: "TX", Country: "USA"},
    }

    // Insert seed data into the database
    for _, person := range people {
        err := db.Create(&person).Error
        if err != nil {
            log.Fatalf("Failed to create person: %v", err)
        }
    }

    log.Println("Seed data created successfully!")
}