package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func main() {
	e := echo.New()

	//e.Debug = true

	// error handler
	//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//  return func(c echo.Context) error {

	//    return errors.New("standard error")
	//    // for invalida login
	//    //return echo.NewHTTPError(http.StatusUnauthorized, "please provide valid credential")
	//    // for valida login
	//    // return next(c)
	//  }
	//})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusInternalServerError, "hello world")
	})

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":1323"))
}
