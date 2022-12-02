package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.File("home.html")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
