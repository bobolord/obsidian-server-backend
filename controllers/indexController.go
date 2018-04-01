package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	yaml "gopkg.in/yaml.v2"
)

var db *gorm.DB
var err error

type DbmsConfigStruct struct {
	Dbms     string `yaml:"dbms,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Database string `yaml:"database,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type Config struct {
	DbmsConfig DbmsConfigStruct `yaml:"DbmsConfig,omitempty"`
}

func Main() *gorm.DB {
	var config Config
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
	db, err = gorm.Open(config.DbmsConfig.Dbms, "host="+config.DbmsConfig.Host+" port="+config.DbmsConfig.Port+" user="+config.DbmsConfig.Username+" dbname="+config.DbmsConfig.Database+" password="+config.DbmsConfig.Password)

	if err != nil {
		panic(err)
	} else {
		log.Print("db connection successful")
	}
	return db
}

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
}
