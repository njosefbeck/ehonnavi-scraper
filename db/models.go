package db

import (
	"github.com/jinzhu/gorm"
)

// Book : Gorm Book model
type Book struct {
	gorm.Model
	URL   string `gorm:"type:varchar(255);unique;not null"`
	Title string `gorm:"type:varchar(255);not null"`
	Age   string `gorm:"type:varchar(10);not null"`
	IsNew bool   `gorm:"-;default:false"`
}
