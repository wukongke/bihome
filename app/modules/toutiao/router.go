package toutiao

import (
	. "work-codes/bihome/app/modules/toutiao/controllers"

	"github.com/labstack/echo"
)

// 币头条模块
func InitRoute(api *echo.Group) {

	api.GET("/users", UserController.List)
	api.POST("/users", UserController.Create)
	api.POST("/demo", UserController.Demo)
}
