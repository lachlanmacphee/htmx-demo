package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	FirstName string
	LastName string
	StreetNumber string
	StreetName string
	Suburb string
	City string
	State string
	Country string
}