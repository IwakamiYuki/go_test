package controller

import (
  "net/http"
	"github.com/labstack/echo"
)
func Hello(c echo.Context) error {
    data := struct {
			ServiceInfo
      Content string
    } {
			ServiceInfo: serviceInfo,
      Content: "おはよう",
    }

    return c.Render(http.StatusOK, "hello.html", data)
}


