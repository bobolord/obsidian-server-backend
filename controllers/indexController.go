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
	db, err = gorm.Open(Config.DbmsConfig.Dbms, "host="+Config.DbmsConfig.Host+" port="+Config.DbmsConfig.Port+" user="+Config.DbmsConfig.Username+" dbname="+Config.DbmsConfig.Database+" password="+Config.DbmsConfig.Password)

	if err != nil {
		panic(err)
	} else {
		log.Print("db connection successful")
	}
	return db
}

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}
