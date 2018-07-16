package controllers

import (
	"log"
	"net/http"

	"github.com/bobolord/obsidian-server-backend/utilities"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Main() *gorm.DB {
	db, err = gorm.Open(utilities.Config.DbmsConfig.Dbms, "host="+utilities.Config.DbmsConfig.Host+" port="+utilities.Config.DbmsConfig.Port+" user="+utilities.Config.DbmsConfig.Username+" dbname="+utilities.Config.DbmsConfig.Database+" password="+utilities.Config.DbmsConfig.Password)

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
