package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type GoogleResponse struct {
}

func GetMovieList() string {
	resp, err := http.Get("https://in.bookmyshow.com/serv/getData?cmd=QUICKBOOK&type=MT&getSeenData=1&getRecommendedData=1&_=1525030467000")
	if err == nil {
		timestamp := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)
		fmt.Println(timestamp)
		// 1525030467000
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		bodyString := string(body)

		if err != nil {
			log.Fatal(err)
		}
		// fmt.Print(bodyString)
		return string(bodyString)
	} else {
		panic(err)
	}
}
