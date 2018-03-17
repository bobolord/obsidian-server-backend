package controllers

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var hmacSampleSecret = []byte("dota")

type loginCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type userTable struct {
	Email    string `gorm:"column:email"`
	Password []byte `gorm:"column:password_hash"`
}

func CheckUserLogin(c *gin.Context) {
	var loginCmd loginCommand
	c.BindJSON(&loginCmd)
	var user userTable

	if !db.Table("users").Where("email = ?", loginCmd.Email).First(&user).RecordNotFound() {
		var success = bcrypt.CompareHashAndPassword(user.Password, []byte(loginCmd.Password))
		if success != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Wrong password", "errorIn": "password"})
		} else {
			c.JSON(http.StatusOK, "Succesfully logged in")
		}
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"foo": "sreedeep",
			"nbf": time.Now().Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		if err == nil {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:     "JWT-TOKEN",
				Value:    tokenString,
				MaxAge:   100000,
				Path:     "/",
				Domain:   "127.0.0.1",
				Secure:   false,
				HttpOnly: false})
			c.JSON(200, gin.H{"message": "Entered email ID isn't registered", "errorIn": "email"})
		}

		// c.JSON(403, gin.H{"message": "Entered email ID isn't registered", "errorIn": "email"})
	}

}

func RegisterUser(c *gin.Context) {
	var loginCmd loginCommand
	c.BindJSON(&loginCmd)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginCmd.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print("gg", err)
	} else {
		user := userTable{Email: loginCmd.Email, Password: hashedPassword}
		if db.Debug().Table("users").Create(&user).Error != nil {
			c.JSON(403, gin.H{"message": "Entered email ID is already registered", "errorIn": "email"})
			return
		}

		c.JSON(http.StatusOK, "succesfully added user")
	}
}

func createToken() {

}

func refreshToken() {

}
