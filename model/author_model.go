package model

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	ID uint
	AuthorName string
	DatePosted string
}
