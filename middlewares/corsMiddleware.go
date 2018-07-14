package middlewares

import (
	"fmt"
	"net/http"

	"github.com/bobolord/obsidian-server-backend/utilities"
)

func CORSMiddleware(h http.Handler) http.Handler {
	fmt.Println("asddadas")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, valid_origin := range utilities.Config.AppConfig.AllowedOrigins {
			if valid_origin == r.Header.Get("Origin") {
				w.Header().Set("Access-Control-Allow-Origin", valid_origin)
			}
		}
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if r.Method == "OPTIONS" {
			http.Error(w, "Please send a request body", 203)
			return
		}
	})
}
