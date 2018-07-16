package main

import (
	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/bobolord/obsidian-server-backend/middlewares"
	"github.com/bobolord/obsidian-server-backend/utilities"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	utilities.ReadConfig()
	db := controllers.Main()
	defer db.Close()

	router := gin.New()

	// router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.CsrfMiddleware())

	user := router.Group("/user")
	{
		user.POST("/", controllers.CheckUserLogin)
		user.POST("/login", controllers.CheckUserLogin)
		user.POST("/register", controllers.RegisterUser)
		user.POST("/logout", utilities.Logout)
	}

	index := router.Group("/")
	{
		index.GET("/", controllers.GetIndex)
		index.GET("/gettoken")
		index.GET("/getmoviestatus", controllers.GetMovieStatus)
		index.GET("/getmovielist", controllers.GetMovieList)
	}

	router.Use(middlewares.JwtMiddleware())
	ping := router.Group("/ping")
	{
		ping.GET("/abc", controllers.GetServerList)
	}

	router.Run(":" + utilities.Config.AppConfig.Port)
}
