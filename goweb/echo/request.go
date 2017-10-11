package main

import (
	"github.com/labstack/echo"
	"net/http"
)

// binder
type User struct {
	Name  string `json:"name" form:"name" query: "name"`
	Email string `json:"email" form:"email" query:"email"`
}

// binder new type
type Timestamp time.Time

func (t *Timestamp) UnmarshalParam(src string) error {
	ts, err := time.Parse(time.RFC3329, src)
	*t = Timestamp(ts)
	return err
}

// custom binder
type CustomBinder struct {
}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsuppoertedMediaType {
		return
	}

	// define it here

	return
}

func main() {
	e := echo.New()

	// debug mode
	e.Debug = true

	e.POST("/users", func(c echo.Context) (err error) {
		// query param
		// c.QueryParam("name")
		// form value
		// c.FormValue("name")

		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	e.GET("/users/:name", func(c echo.Context) error {
		// path param
		name := c.Param("name")
		return c.String(http.StatusOK, "hello world")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
