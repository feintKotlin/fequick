package feint

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"feint/fequick/component"
	"encoding/json"
	"fmt"
	"feint/fequick/logger"
	"os"
	"strings"
	"strconv"
)

type feintGlobe struct {
	routeMap map[string]component.FController
	router   *httprouter.Router
	fConfig  feintConfig
}

type feintConfig struct {
	Server   serverConfig `json:"server"`
	Database databaseConfig
}

type serverConfig struct {
	Port    int    `json:"port"`
	AppName string `json:"app_name"`
	Host    string `json:"host"`
}

type databaseConfig struct {
	Driver   string
	Url      string
	Username string
	Password string
}

var fg feintGlobe

/**
framework的启动应该执行以下几步流程：
1、加载配置文件
2、加载路由
 */
func Run() {
	initBeforeRun()

	addr:=strings.Join([]string{fg.fConfig.Server.Host, ":",
		 strconv.Itoa(fg.fConfig.Server.Port)},"")

	logger.LogI("Run", fg.fConfig.Server.AppName,"On", addr)

	http.ListenAndServe(addr, fg.router)

}

func Route(path string, handle component.FController) {
	if fg.routeMap == nil {
		fg.routeMap = make(map[string]component.FController)
	}

	fg.routeMap[path] = handle
}

func initBeforeRun() {
	loadConfig()
	loadRouter()
}

func loadRouter() {
	fg.router = httprouter.New()

	for key, value := range fg.routeMap {
		httpFunc := func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			inf := value.Execute(params)
			infJson, err := json.Marshal(inf)
			if err == nil {
				fmt.Fprint(w, string(infJson))
			} else {
				//服务端异常
			}
		}
		fg.router.GET(key, httpFunc)
	}
	logger.LogI("Load Router Success")
}

func loadConfig() {
	fg.fConfig = feintConfig{
		Server: serverConfig{
			Port: 8080,
			Host: "127.0.0.1",
			AppName: "FeintApp",
		},
	}

	file, err := os.Open("app.json")
	defer file.Close()
	if err != nil {
		//TODO change to log err
		logger.LogI(err.Error())
		return
	}

	encoder := json.NewDecoder(file)
	err = encoder.Decode(&fg.fConfig)
	if err != nil {
		//TODO change to log err
		logger.LogI(err.Error())
		return
	}

	logger.LogI("Load Config Success")
}
