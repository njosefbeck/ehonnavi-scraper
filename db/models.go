package db

import (
	"github.com/jinzhu/gorm"
	// Required by gorm
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

// Book : Gorm Book model
type Book struct {
	gorm.model
	URL 	string ``
	Title string ``
	Age		string ``
}