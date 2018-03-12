package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Main() {
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=sreedeep dbname=sreedeep password=postgres123")
	// defer db.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
	}
}

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}
