package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello, world")
}

func main() {
	e := echo.New()

	// debug mode
	e.Debug = true

	e.GET("/hello", hello)
	e.Any("/hello", hello)
	e.Match([]string{"GET", "POST"}, "/hello", hello)

	// route group
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string) bool {
		if username == "joe" && password == "secret" {
			return true
		}
		return false
	}))

	// route name
	route := e.POST("/users", func(c echo.Context) error {
		return c.String(200, "hello world")
	})
	route.Name = "create-user"

	e.GET("/users/:id", func(c echo.Context) error {
		data, err := json.MarshalIndent(e.Routes(), "", " ")
		if err != nil {
			return err
		}
		ioutil.WriteFile("routes.json", data, 0644)
		return c.String(200, "hello world")
	}).Name = "get-name"

	// list routes

	e.Logger.Fatal(e.Start(":1323"))
}
