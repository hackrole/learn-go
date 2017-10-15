package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)

	router.Run()
}
