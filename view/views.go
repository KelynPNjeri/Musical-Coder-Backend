package view

import (
	"fmt"
	m "github.com/KelynPNjeri/MC_Backend/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
var DB *gorm.DB



func OpenDBConnection() error {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=blog_backend password=Kelyn@1998")
	if err != nil {
		return err
	}
	return nil
}
func CloseConnection() error {
	return DB.Close()
}

func CreatePost(c *gin.Context) {
	var blog m.Blog
	c.BindJSON(&blog)
	DB.Create(&blog)
	c.JSON(200, blog)
}
func GetAllPosts(c *gin.Context) {
	var posts []m.Blog
	if err := DB.Find(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, posts)
	}
}
func GetSpecificPost (c *gin.Context) {
	id := c.Params.ByName("id")
	var post m.Blog
	if err := DB.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}

}
func UpdatePost(c *gin.Context) {

}

func DeletePost(c *gin.Context) {

}
