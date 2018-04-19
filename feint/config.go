package feint

import (
	"feint/fequick/logger"
	"encoding/json"
	"os"
)

// 加载配置文件
func loadConfig() {
	// 设置配置项的默认值
	fg.fConfig = feintConfig{
		Server: serverConfig{
			Port:    8080,
			Host:    "127.0.0.1",
			AppName: "FeintApp",
		},
		Mongo:MongoDbConfig{
			Url:"localhost",
			Enable:false,
		},
	}

	file, err := os.Open("app.json")
	defer file.Close()
	if err != nil {
		logger.LogE(err)
		return
	}

	encoder := json.NewDecoder(file)
	err = encoder.Decode(&fg.fConfig)
	if err != nil {
		logger.LogE(err)
		return
	}
	logger.LogI("Load Config Success")
}

