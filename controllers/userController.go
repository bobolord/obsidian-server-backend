package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type loginCommand struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type userTable struct {
	Email    string `gorm:"column:email"`
	Password []byte `gorm:"column:password_hash"`
}

func CheckUserLogin(c *gin.Context) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=sreedeep dbname=sreedeep password=postgres123")
	defer db.Close()
	if err != nil {
		panic(err)
	}

	var loginCmd loginCommand
	c.BindJSON(&loginCmd)
	var user userTable

	if !db.Table("users").Where("email = ?", loginCmd.Email).First(&user).RecordNotFound() {
		var success = bcrypt.CompareHashAndPassword(user.Password, []byte(loginCmd.Password))
		if success != nil {
			c.JSON(http.StatusOK, "Wrong password")
		} else {
			c.JSON(http.StatusOK, "Succesfully logged in")
		}
	} else {
		c.JSON(http.StatusOK, "User doesn't exist")
	}
}

func RegisterUser(c *gin.Context) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=sreedeep dbname=sreedeep password=postgres123")
	defer db.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
	}

	var loginCmd loginCommand
	c.BindJSON(&loginCmd)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginCmd.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print("gg", err)
	} else {
		user := userTable{Email: loginCmd.Email, Password: hashedPassword}
		db.Debug().Table("users").Create(&user)
		c.JSON(http.StatusOK, "succesfully added user")
	}
}
