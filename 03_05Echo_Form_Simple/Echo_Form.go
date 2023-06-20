package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func saveUser(c echo.Context) error {
	name := c.FormValue("name")
	family := c.FormValue("family")

	return c.String(http.StatusCreated, name+" + "+family+"\n")
}

func main() {
	e := echo.New()

	e.POST("/users", saveUser)

	e.Logger.Fatal(e.Start(":1323"))
}
