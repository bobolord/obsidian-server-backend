package routers

import (
	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/bobolord/obsidian-server-backend/utilities"
	"github.com/gorilla/mux"
)

func AddUserRouter(r *mux.Router) {
	r.HandleFunc("/ping/abc", controllers.GetServerList).Methods("GET")
	r.HandleFunc("/user/login", controllers.CheckUserLogin).Methods("POST")
	r.HandleFunc("/user/logout", utilities.Logout).Methods("POST")
}
