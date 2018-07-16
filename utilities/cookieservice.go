package utilities

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "XSRF-TOKEN",
		Value:    "hello",
		MaxAge:   -1,
		Path:     "/",
		Domain:   Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "JWT-TOKEN",
		Value:    "hello",
		MaxAge:   -1,
		Path:     "/",
		Domain:   Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
}
