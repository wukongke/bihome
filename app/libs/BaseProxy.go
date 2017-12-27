package libs

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var BaseModelVO *mgo.Collection

func init() {
	model := &BaseModel{}
	BaseModelVO = model.InitModel()
}

type BaseProxy struct{}

func (proxy *BaseProxy) Insert(data ...interface{}) error {
	err := BaseModelVO.Insert(data...)
	return err
}

func (proxy *BaseProxy) FindAll(query interface{}) (rows []*bson.M, err error) {
	err = BaseModelVO.Find(query).All(&rows)
	return
}
func (proxy *BaseProxy) FindOne(query interface{}) (*bson.M, error) {
	var row bson.M
	err := BaseModelVO.Find(query).One(&row)
	return &row, err
}
func (proxy *BaseProxy) FindById(id bson.ObjectId) (*bson.M, error) {
	var row bson.M
	err := BaseModelVO.FindId(id).One(&row)
	return &row, err
}
func (proxy *BaseProxy) Update(selector bson.M, update bson.M) error {
	err := BaseModelVO.Update(selector, update)
	return err
}
