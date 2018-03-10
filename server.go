package main

import (
	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	index := router.Group("/")
	{
		index.GET("/", controllers.GetIndex)
	}

	router.Run(":3001")
}
