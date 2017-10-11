package main

import (
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

func main() {
	e := echo.New()

	// debug mode
	e.Debug = true

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	})

	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.Foo()
		cc.Bar()
		return cc.String(200, "ok")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
