package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// global middleware
	r.Use(gin.Logger())

	// recovery mode
	r.Use(gin.Recovery())

	// per router middleware
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// auth group
	auth := r.Group("/")
	// per group middleware
	auth.Use(AuthRequired())
	{
		auth.POST("/login", loginEndpoint)
		auth.POST("/submit", submitEndpoint)
		auth.POST("/read", readEndpoint)

		// nest group
		testing := auth.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	r.Run(":8080")
}
