package model

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	ID uint
	Title string
	Description string
	DatePosted string
}
