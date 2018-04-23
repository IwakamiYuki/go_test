package main

import (
	"./route"
)

const SERVER_PORT string = ":9999";

func main() {
	e := route.Init()

//	e.File("/", "public/views/index.html")
/*
	e.GET("/", handler.Index())

	e.GET("/users/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, name)
	})

	e.GET("/cookie/set/:name", handler.WriteCookie)
	e.GET("/cookie/get", handler.ReadCookie)
	*/



  e.Logger.Fatal(e.Start(SERVER_PORT))
}

