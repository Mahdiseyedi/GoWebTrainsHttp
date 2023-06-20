package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Family string `json:"family" xml:"family" form:"family" query:"family"`
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.XML(http.StatusCreated, u)
}

func main() {
	e := echo.New()

	e.POST("/users", saveUser)

	e.Logger.Fatal(e.Start(":1323"))
}
