# fequick
a restful framework for quick dev

### 使用事例

#### main.go

```go

import (
	"github.com/feintKotlin/fequick/feint"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Addr Address `json:"addr"`
}

type Address struct {
	City string `json:"city"`
	Town string `json:"town"`
}

type MyController struct {
	feint.FController
}

func (controller MyController) Execute(req feint.FRequest) interface{} {

	if data,err:=req.Data();err==nil{
		collection:=feint.GetCollection("user")
		err=collection.Insert(data)
		if err!=nil{
			return errors.New(fmt.Sprintf("MongoDB Insert: %s",err.Error()))
		}

		return "success"
	}else {
		return err
	}
}

type GetUserController struct {
	feint.FController
}

func (controller GetUserController) Execute(req feint.FRequest) interface{} {

	if _,err:=req.Data();err==nil{
		name:=req.Get("name")

		collection:=feint.GetCollection("user")

		user:=User{}

		collection.Find(bson.M{"name":name}).One(&user)

		return user

	}else{
		return err
	}

}
func init() {
	feint.Route("POST@/user", MyController{})
	feint.Route("GET@/user/:name", GetUserController{})
}

func main() {
	feint.Run()
}


```
app.json
```json

{
  "server": {
    "host": "127.0.0.1",
    "port": 8089,
    "appName": "testApp"
  },
  "mongo":{
    "enable":true,
    "url":"localhost",
    "db":"feint"
  }
}

```

### 待办列表
1. 添加对文件上传的支持

