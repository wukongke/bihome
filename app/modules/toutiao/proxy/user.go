package proxy

import (
	"gopkg.in/mgo.v2/bson"

	"work-codes/bihome/app/common"
	"work-codes/bihome/app/modules/toutiao/models"
)

var UserProxy *userProxy

type userProxy struct {
}

func (proxy *userProxy) Insert(data ...interface{}) error {
	err := models.UserVO.Insert(data...)
	return err
}
func (proxy *userProxy) FindOne(query interface{}) (*bson.M, error) {
	var row bson.M
	err := models.UserVO.Find(query).One(&row)
	return &row, err
}
func (proxy *userProxy) List(filter interface{}, page int, limit int, sort []string) (*map[string]interface{}, error) {
	var rows []map[string]interface{}
	skip := (page - 1) * limit
	err := models.UserVO.Find(filter).Sort(sort...).Skip(skip).Limit(limit).All(&rows)

	total, _ := models.UserVO.Find(filter).Count()
	result := common.PageList(total, page, limit, rows)
	return result, err
}
