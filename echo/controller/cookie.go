package controller

import (
	"net/http"
	"fmt"
	"time"
	"github.com/labstack/echo"
)

func WriteCookie(c echo.Context) error {
	name := c.Param("name")
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = name
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "write a cookie: " + name)
}


func ReadCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie: " + cookie.Value)
}


