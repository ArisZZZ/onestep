package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arisZZZ/onestep/internal/database"
	"github.com/gin-gonic/gin"
)

func RunServer(){
	database.Connect()
	log.Println("🚀: Run server")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})

	r.POST("/user", func(c *gin.Context) {
		fmt.Println("进入啦")
		database.DB.CreateUser(c, "test")
	  })


	r.Run()

}