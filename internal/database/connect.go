package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DataInfo struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Dbname   string `json:"dbname"`
}

var DB *Queries

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

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB = New(db)

	log.Println("✨: Connect database success")
}
