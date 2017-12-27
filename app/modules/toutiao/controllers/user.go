package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"work-codes/bihome/app/common"
	"work-codes/bihome/app/modules/toutiao/proxy"
)

var UserController *userController

type userController struct {
	echo.Context
}

func (c *userController) Demo(ctx echo.Context) error {
	type requestBody struct {
		Name string `json:"name"`
	}
	body := &requestBody{}
	_ = ctx.Bind(body)
	res := echo.Map{
		"user": body,
	}
	return ctx.JSON(http.StatusOK, res)
}

// List 用户列表
func (c *userController) List(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil {
		limit = 15
	}
	filter := bson.M{}
	if search != "" {
		filter["name"] = search
	}
	sort := []string{"-updatedAt"}
	result, _ := proxy.UserProxy.List(filter, page, limit, sort)
	return common.ResJSON(ctx, 0, "操作成功", result)
}

// Create创建用户
func (c *userController) Create(ctx echo.Context) error {
	var body map[string]interface{}
	if err := ctx.Bind(&body); err != nil {
		fmt.Println(err)
		return common.ErrJSON(ctx, -1, "参数错误")
	}
	fmt.Println("body", body)
	if body["name"] == "" {
		return common.ErrJSON(ctx, -1, "用户名称不能为空")
	}
	fmt.Println("body: ")
	row, err := proxy.UserProxy.FindOne(bson.M{"name": body["name"]})
	if err == nil && row != nil {
		return common.ErrJSON(ctx, -1, "用户名称已存在")
	}
	body["status"] = 1
	body["createdAt"] = time.Now().Unix()
	body["updatedAt"] = time.Now().Unix()
	if err := proxy.UserProxy.Insert(body); err != nil {
		return common.ErrJSON(ctx, -1, "创建失败")
	}
	return common.ResJSON(ctx, 0, "操作成功", echo.Map{})
}
