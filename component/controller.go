package component

import "github.com/julienschmidt/httprouter"

type FController interface {
	Execute(params httprouter.Params) interface{}
}



