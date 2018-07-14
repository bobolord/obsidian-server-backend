package middlewares

import (
	"fmt"
	"net/http"

	"github.com/bobolord/obsidian-server-backend/utilities"
)

func CsrfMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		csrfToken, err := r.Cookie("XSRF-TOKEN")
		if err == nil {
			if r.Header.Get("X-Csrf-Token") != "" {
				if csrfToken.Value != r.Header.Get("X-Csrf-Token") {
					fmt.Println("not same csrf", csrfToken.Value, r.Header.Get("X-Csrf-Token"))
					http.Error(w, "Please send a request body", 400)
					return
				}
			} else {
				fmt.Println("x-xsrf nil", csrfToken.Value)
			}
		} else {
			if r.URL.Path != "/gettoken" && err == nil {
				fmt.Println("error with csrf")
				http.Error(w, "Please send a request body", 400)
				return
			}
			cookie := http.Cookie{
				Name:   "XSRF-TOKEN",
				Value:  "hello",
				MaxAge: utilities.Config.AppConfig.CsrfTokenExpiry,
				Path:   "/",
				// Domain:   utilities.GetDomain(w, r),
				Secure:   false,
				HttpOnly: false}
			http.SetCookie(w, &cookie)
		}
	})
}
