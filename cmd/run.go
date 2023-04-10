package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer(){
	log.Println("🚀: Run server")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	  c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	  })
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}