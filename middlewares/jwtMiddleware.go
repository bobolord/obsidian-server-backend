package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bobolord/obsidian-server-backend/utilities"
	jwt "github.com/dgrijalva/jwt-go"
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

func JwtMiddleware(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("requestURI to Jwtware", r.RequestURI)
		jwt, err := CheckJwtToken(w, r)
		if err == nil {
			return
		} else {
			fmt.Println("jwt error", jwt, err)
			http.Error(w, "Please send a request body", 405)
		}
	})
}

func CreateJwtToken(w http.ResponseWriter, r *http.Request) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":           "sai@gmail.com",
		"timeOfCreation": time.Now().Unix(),
	})
	fmt.Println("new token ", time.Now().Unix())
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", fmt.Errorf("Error creating token")
	} else {
		setJwtCookie(w, r, tokenString)
		refreshTokenCookie, err := r.Cookie("REFRESH-TOKEN")
		if err != nil {
			setRefreshCookie(w, r, tokenString)
		} else {
			refreshTokenString := refreshTokenCookie.Value
			refreshToken := JwtToken{}
			token, err := jwt.ParseWithClaims(refreshTokenString, &refreshToken, func(token *jwt.Token) (interface{}, error) {
				return hmacSecret, nil
			})
			if err != nil {
				return "", fmt.Errorf("Error creating token")
			}
			if token.Valid == false {
				utilities.Logout(w, r)
				return "", fmt.Errorf("Authentication error")
			} else {
			}
		}
		return tokenString, nil
	}
}

func refreshJwtToken(w http.ResponseWriter, r *http.Request) (string, error) {
	refreshTokenCookie, err := r.Cookie("REFRESH-TOKEN")
	if err != nil {
		return "", fmt.Errorf("Authentication error")
	}
	refreshTokenString := refreshTokenCookie.Value
	refreshToken := RefreshToken{}
	refreshToken1, err := jwt.ParseWithClaims(refreshTokenString, &refreshToken, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if refreshToken1.Valid == false {
		utilities.Logout(w, r)
		return "", fmt.Errorf("Authentication error")
	}
	jwtToken1 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":           "sai@gmail.com",
		"timeOfCreation": time.Now().Unix(),
	})
	tokenString, err := jwtToken1.SignedString(hmacSecret)
	fmt.Println("Awdadw", tokenString, err)
	if err == nil {
		setJwtCookie(w, r, tokenString)
	}
	return tokenString, nil
}

func setJwtCookie(w http.ResponseWriter, r *http.Request, tokenString string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "JWT-TOKEN",
		Value:    tokenString,
		MaxAge:   utilities.Config.AppConfig.CsrfTokenExpiry,
		Path:     "/",
		Domain:   utilities.Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
}

func setRefreshCookie(w http.ResponseWriter, r *http.Request, tokenString string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "REFRESH-TOKEN",
		Value:    tokenString,
		MaxAge:   101474836, // 3 years
		Path:     "/",
		Domain:   utilities.Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
}

func CheckJwtToken(w http.ResponseWriter, r *http.Request) (string, error) {
	jwtTokenCookie, err := r.Cookie("JWT-TOKEN")
	if err != nil {
		return "", fmt.Errorf("Authentication error")
	}
	jwtTokenString := jwtTokenCookie.Value
	jwtToken := JwtToken{}
	token, err := jwt.ParseWithClaims(jwtTokenString, &jwtToken, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if token.Valid == false {
		utilities.Logout(w, r)
		return "", fmt.Errorf("Authentication error")
	}
	fmt.Println("claims ", jwtToken.TimeOfCreation)
	fmt.Println("claims ", time.Now().Unix())
	if time.Now().Unix()-jwtToken.TimeOfCreation > 10 {
		fmt.Println("New token issued")
		refreshJwtToken(w, r)
	}
	return jwtTokenCookie.Value, nil
}
