package common

import "github.com/labstack/echo"

type Map map[string]interface{}

type Result struct {
	Code int64       `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func ResJSON(ctx echo.Context, code int64, msg string, data interface{}) error {
	res := &Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	return ctx.JSON(200, res)
}

func ErrJSON(ctx echo.Context, code int64, msg string) error {
	res := &Result{
		Code: code,
		Msg:  msg,
		Data: echo.Map{},
	}
	return ctx.JSON(500, res)
}
