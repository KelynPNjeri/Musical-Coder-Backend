package model

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DatePosted string `json:"date_posted"`
}
