package feint

import "gopkg.in/mgo.v2/bson"

// 数据库分页的工具类
func PageUtil(count int, page int, pageNum int) (skip int, limit int) {
	if pageNum == 0 {
		pageNum = 10
	}

	if page*pageNum > count {
		pages := count / pageNum
		skip = pages * pageNum
		limit = count - skip
	} else {
		skip = (page - 1) * pageNum
		limit = pageNum
	}
	return
}


type Count struct {
	Name string
	Count int
}

func PageCountUtil(countName string,countFunc func() int)  (count int){
	countCollection:=GetCollection("count")
	countObj:=Count{}

	countCollection.Find(bson.M{"name":countName}).One(&countObj)

	if countObj.Name=="" {
		count = countFunc()
		countCollection.Insert(Count{Name:countName, Count:count})
	} else {
		count=countObj.Count
	}

	return
}
