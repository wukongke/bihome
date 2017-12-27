package db

import (
	"fmt"
	"work-codes/bihome/app/config"

	"gopkg.in/mgo.v2"
)

func init() {
	fmt.Println("MGO_URL: ", config.DBConfig.Url)
}

// const (
// 	MONGO_URL = "mongodb://127.0.0.1:27017"
// )

func mgoSession() *mgo.Session {
	// 用环境变量
	session, err := mgo.Dial(config.DBConfig.Url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func MgoDB(dbname string) *mgo.Database {
	db := mgoSession().DB(dbname)
	return db
}

func MgoClose() {
	mgoSession().Close()
}
