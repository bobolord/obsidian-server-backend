package main

import (
	"fmt"

	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=sreedeep dbname=sreedeep password=postgres123")
	defer db.Close()

	if err != nil {
		fmt.Println("error connecting to postgresql", err)
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())
	index := router.Group("/")
	{
		index.GET("/", controllers.GetIndex)
	}

	user := router.Group("/user")
	{
		user.POST("/", controllers.CheckUserLogin)
		user.POST("/login", controllers.CheckUserLogin)
		user.POST("/register", controllers.RegisterUser)
	}

	router.Run(":3001")
}
