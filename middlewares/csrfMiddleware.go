package middlewares

import (
	"fmt"
	"net/http"

	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/bobolord/obsidian-server-backend/services/utilities"
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

				} else {
					fmt.Println("successfully authenticated csrf", csrfToken.Value)
				}
			} else {
				fmt.Println("x-xsrf nil", csrfToken.Value, c.Request.Header["X-Csrf-Token"][0])
			}
		} else if c.Request != nil {
			if c.Request.URL.Path != "/gettoken" {
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
		} else {
			fmt.Println("fail")
		}
		c.Next()
	}
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("requestURI to Jwtware", c.Request.RequestURI)
		err := controllers.CheckJwtToken(c)

		// If there is an error, do not call next.
		if err == nil {
			c.Next()
		} else {
			fmt.Println("jwt error", err)
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
		}

	}
}
