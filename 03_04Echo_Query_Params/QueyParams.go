package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getUser(c echo.Context) error {
	name := c.QueryParam("username")
	family := c.QueryParam("family")

	return c.String(http.StatusOK, name+" + "+family+"\n")
}

func main() {
	e := echo.New()

	e.GET("/users", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}
