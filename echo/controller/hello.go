package controller

import (
  "net/http"
	"github.com/labstack/echo"
)
func Hello(c echo.Context) error {
    return c.Render(http.StatusOK, "hello", "World")
}


