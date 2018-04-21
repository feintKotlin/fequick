package feint

import (

	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/feintKotlin/fequick/logger"
)

const   (
	V = "**value**"
)

type feintGlobe struct {
	routeMap map[string] FController
	router   *httprouter.Router
	fConfig  feintConfig
}

type feintConfig struct {
	Server   serverConfig `json:"server"`
	Database databaseConfig
	Mongo mongoDbConfig `json:"mongo"`
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

type mongoDbConfig struct {
	Url string `json:"url"`
	Enable bool `json:"enable"`
	DB string `json:"db"`
}


var fg feintGlobe

/**
framework的启动应该执行以下几步流程：
1、加载配置文件
2、加载路由
*/
func Run() {
	initBeforeRun()

	addr := strings.Join([]string{fg.fConfig.Server.Host, ":",
		strconv.Itoa(fg.fConfig.Server.Port)}, "")

	logger.LogI("Run", fg.fConfig.Server.AppName, "On", addr)

	http.ListenAndServe(addr, fg.router)

}



func initBeforeRun() {
	loadConfig()
	loadRouter()

	if fg.fConfig.Mongo.Enable {
		loadMongo()
	}
}

