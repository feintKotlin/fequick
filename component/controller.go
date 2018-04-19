package component

import (
	"github.com/julienschmidt/httprouter"
)

type RequestMethod string

const (
	M_JSON = "application/json"
)

type FController interface {
	Execute(params httprouter.Params,obj interface{}) interface{}
}




