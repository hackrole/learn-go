package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// disable color
	gin.DisableConsoleColor()

	// logging to a file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// use the following code if you need to write
	// the log to file and console at the same time
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":8080")
}
