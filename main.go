package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.GET("/about", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ini aku yang telah ditipu golang")
	})

	fmt.Println("server strarted on port 5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
