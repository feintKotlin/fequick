# fequick
a restful framework for quick dev

### 使用事例

#### main.go

```go

type MyController struct {
	component.FController
}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Addr Address `json:"addr"`
}

type Address struct {
	City string `json:"city"`
	Town string `json:"town"`
}

func (controller MyController) Execute(params httprouter.Params, obj interface{}) interface{} {
	collection:=feint.GetCollection("user")

	err:=collection.Insert(obj)

	if err!=nil{
		return "Insert To MongoDb Failed"
	}else{
		return obj;
	}
}

type GetUserController struct {
	component.FController
}

func (controller GetUserController) Execute(params httprouter.Params, obj interface{}) interface{} {
	name:=params.ByName("name")

	collection:=feint.GetCollection("user")

	user:=User{}

	collection.Find(bson.M{"name":name}).One(&user)

	return user
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

