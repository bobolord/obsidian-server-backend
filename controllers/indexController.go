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

type Runs struct {
	Dbms     string `yaml:"dbms,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Database string `yaml:"database,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type Document struct {
	Runs []Runs `yaml:"runs,omitempty"`
}

func Main() *gorm.DB {
	var document Document
	reader, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &document)
	if err := yaml.Unmarshal(buf, &document); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println("aaaaaaaaaaaaaaaa", document.Runs[0].Dbms)
	db, err = gorm.Open(document.Runs[0].Dbms, "host="+document.Runs[0].Host+" port="+document.Runs[0].Port+" user="+document.Runs[0].Username+" dbname="+document.Runs[0].Database+" password="+document.Runs[0].Password)

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
