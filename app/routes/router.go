package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"work-codes/bihome/app/config"
	"work-codes/bihome/app/modules"
)

func init() {
	fmt.Println("ApiPre: ", config.APIConfig.Prefix)
}

func logs(ctx echo.Context) error {
	type user struct {
		Name string `json:"name"`
	}
	u := &user{}
	_ = ctx.Bind(u)
	res := echo.Map{
		"user": u,
	}
	return ctx.JSON(http.StatusOK, res)
}

func InitRoute(app *echo.Echo) {
	// 静态文件
	app.Static("/bihome/public", "public")
	app.File("/*", "public/index.html")

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	// module路由注册
	api := app.Group(config.APIConfig.Prefix)
	modules.InitRoute(api)

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, BiHome!")
	})
	api.POST("/logs", logs)
}
