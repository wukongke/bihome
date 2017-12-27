package modules

import (
	"github.com/labstack/echo"

	"work-codes/bihome/app/modules/ico"
	"work-codes/bihome/app/modules/private"
	"work-codes/bihome/app/modules/toutiao"
)

func InitRoute(api *echo.Group) {
	ico.InitRoute(api)
	private.InitRoute(api)
	toutiao.InitRoute(api)
}
