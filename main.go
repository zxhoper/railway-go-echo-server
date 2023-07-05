package main

// Project: /home/hbc/go/src/go_study/go-echo-simple-webapp/echo_basic-001_2023-07-05
import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
