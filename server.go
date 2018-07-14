package main

import (
	"net/http"

	"github.com/bobolord/obsidian-server-backend/routers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/bobolord/obsidian-server-backend/middlewares"
	"github.com/bobolord/obsidian-server-backend/utilities"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	utilities.ReadConfig()
	db := controllers.Main()
	defer db.Close()
	r := mux.NewRouter()
	// user := router.Group("/user")
	// {
	// 	user.POST("/", controllers.CheckUserLogin)
	// 	user.POST("/login", controllers.CheckUserLogin)
	// 	user.POST("/register", controllers.RegisterUser)
	// 	user.POST("/logout", utilities.Logout)
	// }

	// index := router.Group("/")
	// {
	// 	index.GET("/", controllers.GetIndex)
	// 	index.GET("/gettoken")
	// 	index.GET("/getmoviestatus", controllers.GetMovieStatus)
	// 	index.GET("/getmovielist", controllers.GetMovieList)
	// }

	routers.AddPingRouter(r)
	routers.AddIndexRouter(r)
	routers.AddMovieStatusRouter(r)
	routers.AddUserRouter(r)
	n := negroni.Classic()
	recovery := negroni.NewRecovery()
	n.Use(recovery)
	n.UseHandler(middlewares.CsrfMiddleware(r))
	n.UseHandler(middlewares.CORSMiddleware(r))
	// n.UseHandler(middlewares.JwtMiddleware(r))
	n.UseHandler(r)

	// alice1 := alice.New(middlewares.CORSMiddleware, middlewares.CsrfMiddleware, middlewares.JwtMiddleware).Then(n)
	http.ListenAndServe(":"+utilities.Config.AppConfig.Port, n)
}
