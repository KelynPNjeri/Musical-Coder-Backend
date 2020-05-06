package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBConnection() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=blog_backend password=Kelyn@1998")
	if err != nil {
		fmt.Println("Database connection failed.")
	}
	defer db.Close()
}
