package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/bobolord/obsidian-server-backend/middlewares"
	"github.com/bobolord/obsidian-server-backend/utilities"
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

func CheckUserLogin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginCmd loginCommand
	err := decoder.Decode(&loginCmd)
	if err != nil {
		panic(err)
	}

	var user userTable

	if !db.Table("users").Where("email = ?", loginCmd.Email).First(&user).RecordNotFound() {
		var success = bcrypt.CompareHashAndPassword(user.Password, []byte(loginCmd.Password))
		if success != nil {
			json1 := simplejson.New()
			json1.Set("message", "Wrong password")
			json1.Set("errorIn", "Wrong password")
			data, err := json.Marshal(json1)
			if err != nil {
				panic(err)
			}
			http.Error(w, string(data), 400)
		} else {
			middlewares.CreateJwtToken(w, r)
			w.WriteHeader(http.StatusOK)
			// w.Write([]byte("Succesfully logged in"))
		}
	} else {
		responseJson := simplejson.New()
		responseJson.Set("message", "Entered email ID isn't registered")
		responseJson.Set("errorIn", "email")
		jData, err := json.Marshal(responseJson)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jData)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var loginCmd loginCommand
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginCmd)
	if err != nil {
		panic(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(loginCmd.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print("gg", err)
	} else {
		user := userTable{Email: loginCmd.Email, Password: hashedPassword}
		if db.Debug().Table("users").Create(&user).Error != nil {
			responseJson := simplejson.New()
			responseJson.Set("message", "Entered email ID is already registered")
			responseJson.Set("errorIn", "email")
			jData, err := json.Marshal(responseJson)
			if err != nil {
				panic(err)
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(jData)
			return
		}
		// loginCmd.Email = "allspark2020@gmail.com"
		utilities.NewRegistration(loginCmd.Email)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("succesfully added user"))
	}
}
