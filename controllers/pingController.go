package controllers

import (
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

func GetServerList(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("foo", "bar")
	responseBody, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}
	w.Write(responseBody)
}
