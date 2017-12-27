package ico

import (
	"net/http"

	"github.com/labstack/echo"
)

// ICO模块
func InitRoute(api *echo.Group) {

	api.GET("/ico/privates", privates)
}

func privates(ctx echo.Context) error {
	privates := []echo.Map{}
	res := echo.Map{
		"privates": privates,
	}
	return ctx.JSON(http.StatusOK, res)
}
