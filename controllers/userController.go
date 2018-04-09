package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bobolord/obsidian-server-backend/services/utilities"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	yaml "gopkg.in/yaml.v2"
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

func ReadConfig() {
	reader, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &utilities.Config)
	if err := yaml.Unmarshal(buf, &utilities.Config); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func CheckUserLogin(c *gin.Context) {
	var loginCmd loginCommand
	c.BindJSON(&loginCmd)
	var user userTable

	if !db.Table("users").Where("email = ?", loginCmd.Email).First(&user).RecordNotFound() {
		var success = bcrypt.CompareHashAndPassword(user.Password, []byte(loginCmd.Password))
		if success != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong password", "errorIn": "password"})
		} else {
			c.JSON(http.StatusOK, "Succesfully logged in")
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Entered email ID isn't registered", "errorIn": "email"})
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
		loginCmd.Email = "allspark2020@gmail.com"
		utilities.NewRegistration(loginCmd.Email)
		c.JSON(http.StatusOK, "succesfully added user")
	}
}

func Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "XSRF-TOKEN",
		Value:    "hello",
		MaxAge:   -1,
		Path:     "/",
		Domain:   utilities.Config.AppConfig.Domain,
		Secure:   false,
		HttpOnly: false})
}

func createJwtToken() {

}

func refreshJwtToken() {

}

func CheckJwtToken(c *gin.Context) error {
	return fmt.Errorf("Authentication error")

}
