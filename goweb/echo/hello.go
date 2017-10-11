package main

import (
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()

	// root level middelware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// group level middelware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// route level middelware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, world")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})

	e.GET("/show", show)

	e.POST("/save", save)
	e.POST("/save2", save2)

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != -nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func save2(c echo.Context) error {
	name := c.FormValue("name")
	avater, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avater.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avater.Filename)
	if err != nil {
		return err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
