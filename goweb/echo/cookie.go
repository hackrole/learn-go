package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all cookie")
}

func main() {
	e := echo.New()

	// debug mode
	e.Debug = true

	e.GET("/write", writeCookie)
	e.GET("/read", readCookie)
	e.GET("readall", readAllCookies)

	e.Logger.Fatal(e.Start(":1323"))
}
