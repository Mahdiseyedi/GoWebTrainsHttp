package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusCreated, id+"\n")
}

func main() {
	e := echo.New()

	e.GET("/:id", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}
