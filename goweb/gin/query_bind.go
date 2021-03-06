package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:name`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.BindQuery(&person) == nil {
		log.Println("==== only bind by query string ====")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "success")
}
