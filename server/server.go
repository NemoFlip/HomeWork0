package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	err := router.Run("127.0.0.1:8080")
	if err != nil {
		log.Printf("Error of listening and serving: %s", err.Error())
	}
}
