package mainpackage main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "this is post route")
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is get route")
	})
	e.PUT("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "this is put route")
	})
	e.PATCH("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "this is patch route")
	})

	e.Logger.Fatal(e.Start(":1323"))
}


import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "this is post route")
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is get route")
	})
	e.PUT("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "this is put route")
	})
	e.PATCH("/", func(c echo.Context) error {
		return c.String(http.StatusCreated, "this is patch route")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
