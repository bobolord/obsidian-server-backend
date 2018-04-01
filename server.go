package main

import (
	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/bobolord/obsidian-server-backend/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	controllers.ReadConfig()
	db := controllers.Main()
	defer db.Close()

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.CsrfMiddleware())

	user := router.Group("/user")
	{
		user.POST("/", controllers.CheckUserLogin)
		user.POST("/login", controllers.CheckUserLogin)
		user.POST("/register", controllers.RegisterUser)
		user.POST("/logout", controllers.Logout)
	}

	index := router.Group("/")
	{
		index.GET("/", controllers.GetIndex)
		index.GET("/gettoken")
	}

	router.Use(middlewares.JwtMiddleware())

	ping := router.Group("/ping")
	{
		ping.GET("/abc", controllers.GetServerList)
	}

	router.Run(":" + controllers.Config.AppConfig.Port)
}
