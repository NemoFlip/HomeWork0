package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Ok")
	})

	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Printf("Error of listening and serving: %s", err.Error())
	}

}
