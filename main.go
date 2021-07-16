package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func mainHandler(c echo.Context) error {
	fmt.Println("mainHandler çağrıldı")
	return c.String(http.StatusOK, "Main endpointe get isteği yapıldı.")
}

//http://localhost:8088/user/data?username=yenilikci&name=melih&surname=celik
func getUser(c echo.Context) error {
	//param
	dataType := c.Param("data")
	//query
	username := c.QueryParam("username")
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Username: %s, Name: %s, Surname: %s", username, name, surname))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"username": username,
			"name":     name,
			"surname":  surname,
		})
	}

	return c.String(http.StatusBadRequest, "Yalnızca 'json' veya 'string' parametreleri kullanılabilir")
}

type User struct {
	Username string "json:'username'"
	Name     string "json:'name'"
	Surname  string "json:'surname'"
}

func addUser(c echo.Context) error {
	user := User{}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	fmt.Println(user)
	return c.String(200, "Başarılı")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Admin endpointindesin!")
}

func setHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")
		if contentType != "application/json" {
			return c.String(http.StatusBadRequest, "Yalnızca application/json tipinde istek atılabilir!")
		}
		return next(c)
	}
}

func main() {
	fmt.Printf("Hello World")

	e := echo.New()
	e.Use(setHeader)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "statusCode=${status}\n",
	}))

	e.GET("/main", mainHandler)

	e.GET("/user/:data", getUser)

	e.POST("/user", addUser)

	//adminGroup := e.Group("/admin", middleware.Logger())
	adminGroup := e.Group("/admin")
	// /admin/help
	adminGroup.GET("/main", mainAdmin)

	e.Start(":8080")
}
