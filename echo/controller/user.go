package controller

import (
	"net/http"
	"github.com/labstack/echo"
)

func User() echo.HandlerFunc {
    return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
				name := c.Param("name")
        return c.String(http.StatusOK, "Hello, " + name + " World")
    }
}

