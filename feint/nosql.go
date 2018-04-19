package feint

import (
	"gopkg.in/mgo.v2"

	"github.com/feintKotlin/fequick/logger"
)

type MongoConn struct {
	DB *mgo.Database
}

var mongoC MongoConn

func loadMongo()  {
	conn,err:=mgo.Dial(fg.fConfig.Mongo.Url)
	if err!=nil{
		logger.LogE(err)
	}

	if db:=fg.fConfig.Mongo.DB;db==""{
		logger.LogE("Database Name Cannot Be Null")
	}else{
		mongoC.DB=conn.DB(db)
	}

	logger.LogI("Load MongoDb Success")
}


func GetCollection(collection string) *mgo.Collection{
	return mongoC.DB.C(collection)
}