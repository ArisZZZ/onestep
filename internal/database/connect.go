package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dbname   string `json:"dbname"`
}

var DB *gorm.DB

func Connect() {
	var dataInfo DataInfo
	jsonData, err := ioutil.ReadFile("config/database.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonData, &dataInfo)
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dataInfo.User, dataInfo.Password, dataInfo.Host, dataInfo.Port, dataInfo.Dbname)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	log.Println("✨: Connect database success")
}
