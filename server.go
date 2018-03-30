package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bobolord/obsidian-server-backend/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	yaml "gopkg.in/yaml.v2"
)

type Config_Struct struct {
	App_Config App_Config_Struct `yaml:"app_Config,omitempty"`
}

type App_Config_Struct struct {
	Allowed_origins string `yaml:"allowed_origins,omitempty"`
	Port            string `yaml:"port,omitempty"`
}

var config Config_Struct

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", config.App_Config.Allowed_origins)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Skip-Interceptor, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func CsrfMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrfToken, err := c.Request.Cookie("XSRF-TOKEN")
		if err == nil {
			if c.Request.Header["X-Csrf-Token"] != nil {
				if csrfToken.Value != c.Request.Header["X-Csrf-Token"][0] {
					fmt.Println(csrfToken.Value, c.Request.Header["X-Csrf-Token"][0])
				} else {
					fmt.Println("successfully authenticated csrf", csrfToken.Value)
				}
			} else {
				fmt.Println("x-xsrf nil", csrfToken.Value, c.Request.Header["X-Csrf-Token"][0])
			}
		} else if c.Request != nil {
			if c.Request.URL.Path != "/gettoken" {
				fmt.Println("false")
			} else {
				fmt.Println("true")
				http.SetCookie(c.Writer, &http.Cookie{
					Name:     "XSRF-TOKEN",
					Value:    "hello",
					MaxAge:   100000,
					Path:     "/",
					Domain:   "127.0.0.1",
					Secure:   false,
					HttpOnly: false})
			}
		} else {
			fmt.Println("fail")
		}
		c.Next()
	}
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Request.Body

		fmt.Println(user)
		c.Next()
	}
}

func main() {
	reader, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &config)
	if err := yaml.Unmarshal(buf, &config); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	db := controllers.Main()
	defer db.Close()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())
	router.Use(CsrfMiddleware())

	index := router.Group("/")
	{
		index.GET("/", controllers.GetIndex)
		index.GET("/gettoken")
	}

	user := router.Group("/user")
	{
		user.POST("/", controllers.CheckUserLogin)
		user.POST("/login", controllers.CheckUserLogin)
		user.POST("/register", controllers.RegisterUser)
	}
	// user.Use(JwtMiddleware())

	router.Run(":" + config.App_Config.Port)
}
