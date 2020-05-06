package model

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	ID uint `json:"id"`
	AuthorName string `json:"author_name"`
	DatePosted string `json:"date_posted"`
}
