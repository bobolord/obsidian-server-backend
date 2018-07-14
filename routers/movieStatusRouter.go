package routers

import (
	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/gorilla/mux"
)

func AddMovieStatusRouter(r *mux.Router) {
	r.HandleFunc("/ping/abc", controllers.GetServerList).Methods("GET")
}
