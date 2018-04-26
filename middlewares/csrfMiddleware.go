package middlewares

import (
	"fmt"
	"net/http"

	"github.com/bobolord/obsidian-server-backend/utilities"
	"github.com/gin-gonic/gin"
)

func CsrfMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrfToken, err := c.Request.Cookie("XSRF-TOKEN")
		if err == nil {
			if c.Request.Header["X-Csrf-Token"] != nil {
				if csrfToken.Value != c.Request.Header["X-Csrf-Token"][0] {
					fmt.Println("not same csrf", csrfToken.Value, c.Request.Header["X-Csrf-Token"][0])
					c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
					c.Abort()
				}
			} else {
				fmt.Println("x-xsrf nil", csrfToken.Value, c.Request.Header["X-Csrf-Token"][0])
			}
		} else {
			if c.Request.URL.Path != "/gettoken" && c.Request.URL.Path != "/getmoviestatus" {
				fmt.Println("error with csrf")
				c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
				c.Abort()

			} else {
				fmt.Println("true")
				http.SetCookie(c.Writer, &http.Cookie{
					Name:     "XSRF-TOKEN",
					Value:    "hello",
					MaxAge:   utilities.Config.AppConfig.CsrfTokenExpiry,
					Path:     "/",
					Domain:   utilities.Config.AppConfig.Domain,
					Secure:   false,
					HttpOnly: false})
			}
		}
		c.Next()
	}
}
