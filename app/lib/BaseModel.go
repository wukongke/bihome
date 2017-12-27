package lib

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"work-codes/bihome/app/common"
	"work-codes/bihome/app/config"
	"work-codes/bihome/app/db"
)

// var BaseModelVO *mgo.Collection

type BaseModel struct {
	Id        bson.ObjectId `bson:"_id"`
	Status    int           `bson:"status"`
	IsDeleted int           `bson:"isDeleted"`
	CreatedAt int           `bson:"createdAt"`
	UpdatedAt int           `bson:"updatedAt"`
	DeletedAt int           `bson:"deletedAt"`
}

func (model *BaseModel) InitModel() *mgo.Collection {
	baseModelVO := db.MgoDB(config.DBConfig.DbName).C("tb_test")
	return baseModelVO
}

// ToMap struct转成map
func (model *BaseModel) ToMap() common.Map {
	rowMap := common.Map{}
	rowMap["id"] = model.Id
	rowMap["createdAt"] = model.CreatedAt
	rowMap["updatedAt"] = model.UpdatedAt
	return rowMap
}
