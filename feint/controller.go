package feint

import (
	"github.com/julienschmidt/httprouter"
)


type FRequest struct {
	params httprouter.Params
	data interface{}
}

func (req *FRequest) Get(key string) string {
	return req.params.ByName(key)
}

func (req *FRequest) Data() (map[string] interface{},error) {

	switch data:=req.data.(type) {
	case error:
		return nil,data
	case map[string] interface{}:
		return data,nil
	default:
		m:=make(map[string]interface{})
		m[V]=data
		return m,nil
	}
}

type FController interface {
	Execute(req FRequest) interface{}
}





