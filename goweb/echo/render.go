package main

import (
	"github.com/labstack/echo"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// debug mode
	e.Debug = true

	t = &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/hello", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello", "world")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
