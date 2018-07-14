package utilities

import (
	"fmt"
	"net/http"
)

func GetDomain(w http.ResponseWriter, r *http.Request) string {
	var domain string
	parsedUrl := r.URL
	if parsedUrl.Hostname() == "localhost" {
		fmt.Println("Asas" + domain)
		return "localhost"
	}
	domain = parsedUrl.Hostname()
	return domain
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "JWT-TOKEN",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
	http.SetCookie(w, &http.Cookie{
		Name:     "REFRESH-TOKEN",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
}
