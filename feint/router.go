package feint

import (
	"strings"

	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/feintKotlin/fequick/logger"
	"errors"
)

func Route(path string, handle func(request FRequest)interface{}) {
	if fg.routeMap == nil {
		fg.routeMap = make(map[string] func(request FRequest)interface{})
	}

	fg.routeMap[path] = handle

	logger.LogI("Loading Route: map", path,"to",reflect.TypeOf(handle).Name())

}

func loadRouter() {
	fg.router = httprouter.New()
	for key, value := range fg.routeMap {

		keys:=strings.Split(key,"@")
		if len(keys)!=2 {
			logger.LogE("Invalid path", key)
		}

		httpFunc := func(value func(request FRequest)interface{}, path string) (func(w http.ResponseWriter, r *http.Request, params httprouter.Params)){

			return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

				logger.LogI("struct =",reflect.TypeOf(value).Name(),"; path =",path,"; address =",r.RemoteAddr)

				var data interface{}

				if keys[0] == "POST" {
					err := json.NewDecoder(r.Body).Decode(&data)
					if err != nil {
						logger.LogW("Json:",err)
						data = errors.New(fmt.Sprintf("Json Decoder Error: %s",err.Error()))
					}

					r.Body.Close()
				}


				inf := value(FRequest{params, data})

				infJson, err := json.Marshal(inf)
				if err == nil {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					fmt.Fprint(w, string(infJson))
				} else {
					logger.LogW(err)
				}
			}
		}(value, keys[1])

		logger.LogI("Method:",keys[0],"; Path:",keys[1])

		fg.router.Handle(keys[0], keys[1], httpFunc)
	}
	logger.LogI("Load Router Success")
}

