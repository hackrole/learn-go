package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// debug mode
	e.Debug = true

	e.Static("/static", "./")
	// register file
	e.File("/route", "/route.go")

	e.Logger.Fatal(e.Start(":1323"))
}
