package controllers

import (
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/bobolord/obsidian-server-backend/utilities"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func Main() *gorm.DB {
	db, err = gorm.Open(utilities.Config.DbmsConfig.Dbms, "host="+utilities.Config.DbmsConfig.Host+" port="+utilities.Config.DbmsConfig.Port+" user="+utilities.Config.DbmsConfig.Username+" dbname="+utilities.Config.DbmsConfig.Database+" password="+utilities.Config.DbmsConfig.Password)

	if err != nil {
		panic(err)
	} else {
		log.Print("db connection successful")
	}
	return db
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("foo", "index")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	w.Write(payload)
}
