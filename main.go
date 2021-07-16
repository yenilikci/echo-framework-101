package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Main endpointe get isteği yapıldı.")
}

func main() {
	fmt.Printf("Hello World")

	e := echo.New()
	e.GET("/main", mainHandler)

	e.Start(":8088")
}
