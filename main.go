package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func mainHandler(c echo.Context) error {
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

func main() {
	fmt.Printf("Hello World")

	e := echo.New()

	e.GET("/main", mainHandler)

	e.GET("/user/:data", getUser)

	e.POST("/user", addUser)

	e.Start(":8088")
}
