package private

import (
	"github.com/labstack/echo"

	. "work-codes/bihome/app/modules/private/controllers"
)

// 私募模块

func InitRoute(api *echo.Group) {

	api.GET("/categorys", CategoryController.List)
	api.GET("/categorys/:id", CategoryController.Detail)
	api.POST("/categorys", CategoryController.Create)
	api.POST("/categorys/:id", CategoryController.Edit)
}
