package libs

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"work-codes/bihome/app/common"
	"work-codes/bihome/app/config"
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
	baseModelVO := common.DB(config.DBConfig.DbName).C("tb_test")
	return baseModelVO
}
