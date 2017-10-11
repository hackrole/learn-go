package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func main() {
	e := echo.New()

	// debug mode
	e.Debug = true

	e.GET("/", func(c echo.Context) error {
		// string
		return c.String(http.StatusOK, "hello world")
		// html
		return c.HTML(http.StatusOK, "<strong>hello</strong>")
		// return c.HTMLBlob(...)

		// json
		u := &User{
			Name:  "Jon",
			Email: "Jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)

		// stream json
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(u)
		// json pretty
		return c.JSONPretty(http.StatusOK, u, "  ")
		// json blob
		encodedJSON := []byte{}
		return c.JSONBlob(http.StatusOK, encodedJSON)
		// jsonp
		// return c.JSONP(200, "")

		// xml
		return c.XML(http.StatusOK, u)
		// stream xml
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return xml.NewEncoder(c.Response()).Encode(u)
		// xml pretty
		return c.XMLPretty(http.StatusOK, u, " ")

		// send file
		return c.File("<path-to-file>")
		// set attachmetn
		return c.Attachment("<path-to-file>")
		// set inline-file
		return c.Inline("<path-to-file>")

		// send blob
		return c.Blob(https.StatusOK, "text/csv", data)

		// send stream
		f, err := os.Open("<path-to-file>")
		if err != nil {
			return nil
		}
		return c.Stream(https.StatusOK, "image/png", f)

		// noContent
		return c.NoContent(https.StatusOK)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
