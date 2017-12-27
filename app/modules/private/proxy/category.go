package proxy

import (
	"gopkg.in/mgo.v2/bson"

	"work-codes/bihome/app/common"
	"work-codes/bihome/app/modules/private/models"
)

var CategoryProxy *categoryProxy

type categoryProxy struct{}

func (proxy *categoryProxy) Insert(data ...interface{}) error {
	err := models.CategoryVO.Insert(data...)
	return err
}
func (proxy *categoryProxy) FindAll(query interface{}) (rows []*models.Category, err error) {
	err = models.CategoryVO.Find(query).All(&rows)
	return
}
func (proxy *categoryProxy) FindOne(query interface{}) (*models.Category, error) {
	var row models.Category
	err := models.CategoryVO.Find(query).One(&row)
	return &row, err
}
func (proxy *categoryProxy) FindById(id bson.ObjectId) (*models.Category, error) {
	var row models.Category
	err := models.CategoryVO.FindId(id).One(&row)
	return &row, err
}
func (proxy *categoryProxy) Update(selector bson.M, update bson.M) error {
	err := models.CategoryVO.Update(selector, update)
	return err
}
func (proxy *categoryProxy) UpdateAll(selector bson.M, data bson.M) (info interface{}, err error) {
	info, err = models.CategoryVO.UpdateAll(selector, data)
	return
}
func (proxy *categoryProxy) UpdateId(id bson.ObjectId, data interface{}) error {
	err := models.CategoryVO.UpdateId(id, data)
	return err
}
func (proxy *categoryProxy) Upsert(selector interface{}, data interface{}) (info interface{}, err error) {
	info, err = models.CategoryVO.Upsert(selector, data)
	return
}
func (proxy *categoryProxy) BatchRemove(filter interface{}) error {
	err := models.CategoryVO.Remove(filter)
	return err
}
func (proxy *categoryProxy) RemoveId(id interface{}) error {
	err := models.CategoryVO.RemoveId(id)
	return err
}
func (proxy *categoryProxy) List(filter interface{}, page int, limit int, sort []string) (*map[string]interface{}, error) {
	var rows []map[string]interface{}
	skip := (page - 1) * limit
	err := models.CategoryVO.Find(filter).Sort(sort...).Skip(skip).Limit(limit).All(&rows)

	total, _ := models.CategoryVO.Find(filter).Count()
	result := common.PageList(total, page, limit, rows)
	return result, err
}

func (proxy *categoryProxy) Top(filter interface{}, page int, limit int, sort []string) ([]*map[string]interface{}, error) {
	var rows []*map[string]interface{}
	skip := (page - 1) * limit
	err := models.CategoryVO.Find(filter).Sort(sort...).Skip(skip).Limit(limit).All(&rows)
	return rows, err
}
