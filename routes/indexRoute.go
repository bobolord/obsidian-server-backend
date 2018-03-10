package routes

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/gin-gonic/gin"

)

func Routes(route *gin.Engine)
indexRoute = route.Group("/"){
    index.GET("/")
}
