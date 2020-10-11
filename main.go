package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)
type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
}
func home (c echo.Context) error{
	var array [5]map[string]string
	user := User{Name: "Edwin Baeza", Email:"edwinbaeza05@gmail.com"}
	client := new(User)
	client.Name = "Ing. Edwin Baeza"
	client.Email = "edwinbaeza05@gmail.com"
	//return c.String(http.StatusOK, "Hello, World!")
	data := make(map[string]string)
	data["name"] = client.Name
	data["email"] = client.Email
	array[0] = data
	for key, value := range data{
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println("------------------")
	}
	return c.JSON(http.StatusOK, user)
}
func main(){
	fmt.Println("API Started...!")
	e := echo.New()
	e.GET("/", home)
	e.Start(":1323")
}