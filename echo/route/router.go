package route

import(
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"../controller"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/index.html")

	e.GET("/user/:name", controller.User())

	e.GET("/cookie/set/:name", controller.WriteCookie)
	e.GET("/cookie/get", controller.ReadCookie)

	return e
}
