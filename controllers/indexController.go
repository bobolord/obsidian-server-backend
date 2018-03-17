package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Main() *gorm.DB {

	db, err = gorm.Open("postgres", "host=localhost port=5432 user=sreedeep dbname=server_status_app password=postgres123")
	// defer db.Close()
	if err != nil {
		panic(err)
	} else {
		log.Print("failed to connect to database")
	}
	return db
}

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}
