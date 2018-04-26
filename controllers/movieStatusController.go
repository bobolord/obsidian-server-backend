package controllers

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetMovieStatus(c *gin.Context) {
	resp, err := http.Get("https://in.bookmyshow.com/buytickets/avengers-infinity-war-3d-hyderabad/movie-hyd-ET00053419-MT/20180427")
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			movieStatus := strings.Contains(string(body), "/buytickets/prasads-large-screen/cinema-hyd-PRHY-MT/20180427")
			c.JSON(http.StatusOK, gin.H{"status": movieStatus})
		}
	}
}
