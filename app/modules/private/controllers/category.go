package controllers

import (
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"

	"work-codes/bihome/app/common"
	"work-codes/bihome/app/modules/private/proxy"
)

var CategoryController *categoryController

type categoryController struct {
	echo.Context
}

// List 分类列表
func (c *categoryController) List(ctx echo.Context) error {
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
	result, _ := proxy.CategoryProxy.List(filter, page, limit, sort)

	// docs := []map[string]interface{}{}
	// for _, category := range result["docs"].([]map[string]interface{}) {
	// 	docs = append(docs, *proxy.Convert2Book(category))
	// }
	// result["docs"] = docs
	return common.ResJSON(ctx, 0, "操作成功", result)
}

// Detail 获取分类信息
func (c *categoryController) Detail(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return common.ErrJSON(ctx, -1, "id不能为空")
	}
	row, err := proxy.CategoryProxy.FindOne(bson.M{"id": id})
	if err != nil {
		return common.ErrJSON(ctx, -1, "分类不存在")
	}
	// row = proxy.Convert2Category(row)
	return common.ResJSON(ctx, 0, "操作成功", echo.Map{
		"category": row,
	})
}

// Create创建分类
func (c *categoryController) Create(ctx echo.Context) error {
	var body map[string]interface{}
	if err := ctx.Bind(body); err != nil {
		return common.ErrJSON(ctx, -1, "参数错误")
	}
	if body["name"] == "" {
		return common.ErrJSON(ctx, -1, "分类名称不能为空")
	}
	row, err := proxy.CategoryProxy.FindOne(bson.M{"name": body["name"]})
	if err == nil && row != nil {
		return common.ErrJSON(ctx, -1, "分类名称已存在")
	}
	body["status"] = 1
	body["createdAt"] = time.Now().Unix()
	body["updatedAt"] = time.Now().Unix()
	if err := proxy.CategoryProxy.Insert(body); err != nil {
		return common.ErrJSON(ctx, -1, "创建失败")
	}
	return common.ResJSON(ctx, 0, "操作成功", echo.Map{})
}

// 更新分类
func (c *categoryController) Edit(ctx echo.Context) error {
	id := bson.ObjectIdHex(ctx.Param("id"))
	type RequestBody struct {
		Name     string `json:"name"`
		Sequence int    `json:"sequence"`
		ParentID string `json:"parentId"`
	}
	var body RequestBody
	if err := ctx.Bind(body); err != nil {
		return common.ErrJSON(ctx, -1, "参数错误")
	}
	if id == "" {
		return common.ErrJSON(ctx, -1, "分类ID不能为空")
	}
	if err := proxy.CategoryProxy.UpdateId(id, body); err != nil {
		return common.ErrJSON(ctx, -1, "操作失败")
	}
	return common.ResJSON(ctx, 0, "操作成功", echo.Map{})
}

// 删除分类 (如果分类下有标签， 则不可删除)
func (c *categoryController) Delete(ctx echo.Context) error {
	id := bson.ObjectIdHex(ctx.Param("id"))
	// if _, err := TagService.FindOne(bson.M{"cid": id}); err == nil {
	// 	return common.ErrJSON(ctx, -1, "此分类下有话题")
	// }
	if err := proxy.CategoryProxy.RemoveId(id); err != nil {
		return common.ErrJSON(ctx, -1, "删除失败")
	}
	return common.ResJSON(ctx, 0, "操作成功", echo.Map{})
}
