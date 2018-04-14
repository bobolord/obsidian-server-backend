package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bobolord/obsidian-server-backend/utilities"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var hmacSecret = []byte("Obsidian")

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("requestURI to Jwtware", c.Request.RequestURI)
		jwt, err := CheckJwtToken(c)
		if err == nil {
			c.Next()
		} else {
			fmt.Println("jwt error", jwt, err)
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			c.Abort()
		}
	}
}

func CreateJwtToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Unix(),
	})
	tokenString, err := token.SignedString(hmacSecret)
	fmt.Println(tokenString, err)
	if err == nil {
		setJwtCookie(c, tokenString)
	}
}

func refreshJwtToken(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Now().Unix(),
	})
	tokenString, err := token.SignedString(hmacSecret)
	fmt.Println(tokenString, err)
	if err == nil {
		setJwtCookie(c, tokenString)
	}
}

func setJwtCookie(c *gin.Context, tokenString string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "JWT-TOKEN",
		Value:    tokenString,
		MaxAge:   utilities.Config.AppConfig.CsrfTokenExpiry,
		Path:     "/",
		Domain:   utilities.Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
}

func CheckJwtToken(c *gin.Context) (string, error) {
	jwtToken, err := c.Request.Cookie("JWT-TOKEN")
	jwtTokenString := jwtToken.Value
	token, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSecret, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		if err != nil {
			return "", fmt.Errorf("Authentication error")
		} else {
			if tokenTime, ok := claims["nbf"].(int64); ok {
				fmt.Println(time.Now().Unix() - tokenTime)
				if time.Now().Unix()-tokenTime > 0 {
					fmt.Println(time.Now().Unix() - tokenTime)
					refreshJwtToken(c)
				}
			} else {
				fmt.Println("Asas", tokenTime)
			}
		}

	} else {
		fmt.Println(err)
	}
	return jwtToken.Value, nil

}
