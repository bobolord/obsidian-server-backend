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

type JwtToken struct {
	User           string `json:"user"`
	TimeOfCreation int64  `json:"timeOfCreation"`
	jwt.StandardClaims
}
type RefreshToken struct {
	User           string `json:"user"`
	TimeOfCreation int64  `json:"timeOfCreation"`
	jwt.StandardClaims
}

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

func CreateJwtToken(c *gin.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":           "sai@gmail.com",
		"timeOfCreation": time.Now().Unix(),
	})
	fmt.Println("new token ", time.Now().Unix())
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", fmt.Errorf("Error creating token")
	} else {
		setJwtCookie(c, tokenString)
		refreshTokenCookie, err := c.Request.Cookie("REFRESH-TOKEN")
		if err != nil {
			setRefreshCookie(c, tokenString)
		}
		refreshTokenString := refreshTokenCookie.Value
		refreshToken := JwtToken{}
		token, err := jwt.ParseWithClaims(refreshTokenString, &refreshToken, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})
		if token.Valid == false {
			utilities.Logout(c)
			return "", fmt.Errorf("Authentication error")
		} else {
		}
		return tokenString, nil
	}
}

func refreshJwtToken(c *gin.Context) (string, error) {
	refreshTokenCookie, err := c.Request.Cookie("REFRESH-TOKEN")
	if err != nil {
		return "", fmt.Errorf("Authentication error")
	}
	refreshTokenString := refreshTokenCookie.Value
	refreshToken := RefreshToken{}
	refreshToken1, err := jwt.ParseWithClaims(refreshTokenString, &refreshToken, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if refreshToken1.Valid == false {
		utilities.Logout(c)
		return "", fmt.Errorf("Authentication error")
	}
	jwtToken1 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":           "sai@gmail.com",
		"timeOfCreation": time.Now().Unix(),
	})
	tokenString, err := jwtToken1.SignedString(hmacSecret)
	fmt.Println("Awdadw", tokenString, err)
	if err == nil {
		setJwtCookie(c, tokenString)
	}
	return tokenString, nil
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

func setRefreshCookie(c *gin.Context, tokenString string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "REFRESH-TOKEN",
		Value:    tokenString,
		MaxAge:   101474836, // 3 years
		Path:     "/",
		Domain:   utilities.Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
}

func CheckJwtToken(c *gin.Context) (string, error) {
	jwtTokenCookie, err := c.Request.Cookie("JWT-TOKEN")
	if err != nil {
		return "", fmt.Errorf("Authentication error")
	}
	jwtTokenString := jwtTokenCookie.Value
	jwtToken := JwtToken{}
	token, err := jwt.ParseWithClaims(jwtTokenString, &jwtToken, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if token.Valid == false {
		utilities.Logout(c)
		return "", fmt.Errorf("Authentication error")
	} else {
		fmt.Println("claims ", jwtToken.TimeOfCreation)
		fmt.Println("claims ", time.Now().Unix())
		if time.Now().Unix()-jwtToken.TimeOfCreation > 10 {
			fmt.Println("New token issued")
			refreshJwtToken(c)
		}
	}

	return jwtTokenCookie.Value, nil

}
