package route

import(
  "html/template"
  "io"
	"os"
	"net/http"
	"fmt"

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

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("public/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}


func Init() *echo.Echo {
	e := echo.New()

  // ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


 fp, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    panic(err)
  }

  e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Output: fp,
  }))


  renderer := &Template{
    templates: template.Must(template.ParseGlob("public/views/*.html")),
  }
  e.Renderer = renderer

	// カスタムエラーハンドリング
	e.HTTPErrorHandler = customHTTPErrorHandler


	// 静的ページのテスト
	e.File("/", "public/index.html")

	// 引数のテスト
	e.GET("/user/:name", controller.User())

	// クッキーのテスト
	e.GET("/cookie/set/:name", controller.WriteCookie)
	e.GET("/cookie/get", controller.ReadCookie)

	// テンプレートのテスト
  e.GET("/hello", controller.Hello)

	// 静的ファイルの読み込み
	e.Static("/css", "public/css")
	e.Static("/image", "public/images")

	return e
}
