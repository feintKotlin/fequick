package feint

import (
	"strings"

	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/feintKotlin/fequick/logger"
	"github.com/feintKotlin/fequick/component"
)

func Route(path string, handle component.FController) {
	if fg.routeMap == nil {
		fg.routeMap = make(map[string] component.FController)
	}

	fg.routeMap[path] = handle

	logger.LogI("Loading Route: map", path,"to",reflect.ValueOf(handle).Kind().String())

}

func loadRouter() {
	fg.router = httprouter.New()
	for key, value := range fg.routeMap {

		keys:=strings.Split(key,"@")
		if len(keys)!=2 {
			logger.LogE("Invalid path", key)
		}

		httpFunc := func(value component.FController) (func(w http.ResponseWriter, r *http.Request, params httprouter.Params)){

			return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

				var data interface{}

				if keys[0] == "POST" {
					err := json.NewDecoder(r.Body).Decode(&data)
					if err != nil {
						logger.LogW(err)
						data = "InValid Json Struck"
					}

					r.Body.Close()
				}

				inf := value.Execute(params, data)

				infJson, err := json.Marshal(inf)
				if err == nil {
					fmt.Fprint(w, string(infJson))
				} else {
					logger.LogW(err)
				}
			}
		}(value)

		logger.LogI("Method:",keys[0],"; Path:",keys[1])

		fg.router.Handle(keys[0], keys[1], httpFunc)
	}
	logger.LogI("Load Router Success")
}

