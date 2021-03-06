package models

import (
	"work-codes/bihome/app/common"
	"work-codes/bihome/app/config"
	"work-codes/bihome/app/db"
	"work-codes/bihome/app/lib"
)

var UserVO = db.MgoDB(config.DBConfig.DbName).C("tb_user")

// 用户表
type User struct {
	lib.BaseModel
	Name string `bson:"name"`
}

// ToJSON 转成map
func (model *User) ToJSON() common.Map {
	return common.Map{
		"id":        model.Id,
		"name":      model.Name,
		"status":    model.Status,
		"isDeleted": model.IsDeleted,
		"createdAt": model.CreatedAt,
		"updatedAt": model.UpdatedAt,
		"deletedAt": model.DeletedAt,
	}
}
