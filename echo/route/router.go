package route

import(
  "html/template"
  "io"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"../controller"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

func Init() *echo.Echo {
	e := echo.New()

  // ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

  renderer := &Template{
    templates: template.Must(template.ParseGlob("public/views/*.html")),
  }

  e.Renderer = renderer


	e.File("/", "public/index.html")

	e.GET("/user/:name", controller.User())

	e.GET("/cookie/set/:name", controller.WriteCookie)
	e.GET("/cookie/get", controller.ReadCookie)

  e.GET("/hello", controller.Hello)

	return e
}
