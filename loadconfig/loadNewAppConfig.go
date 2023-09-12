package loadconfig

import (
	"encoding/json"
	"os"
	"reflect"
	"time"
	gconfig "xj/xapi-backend/g_config"
	glog "xj/xapi-backend/g_log"
)

// 从环境变量中获取 MySQL 连接信息
// dbHost := os.Getenv("DB_HOST")
// dbPort := os.Getenv("DB_PORT")
// dbUser := os.Getenv("DB_USER")
// dbPassword := os.Getenv("DB_PASSWORD")
// dbName := os.Getenv("DB_NAME")
// if !utils.AreEmptyStrings(dbHost, dbPort, dbUser, dbPassword, dbName) {
// 	gconfig.AppConfig.Database.Host = dbHost
// 	gconfig.AppConfig.Database.Port, _ = strconv.Atoi(dbPort)
// 	gconfig.AppConfig.Database.Dbname = dbName
// 	gconfig.AppConfig.Database.Username = dbUser
// 	gconfig.AppConfig.Database.Password = dbPassword
// } else {
// 	glog.Log.Info("environment dbconfig error!!")
// }

// 构建 MySQL 连接字符串
// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

// 加载App配置数据
func LoadAppConfig() (*gconfig.AppConfiguration, error) {
	filePath := "conf/appconfig.json"
	config := &gconfig.AppConfiguration{}

	// 打开项目配置文件
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	// 解码配置文件内容到结构体
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

func LoadNewAppConfig() {
	filePath := "conf/appconfig.json"
	ticker := time.NewTicker(3 * time.Second) // 每3秒检查一次配置文件
	defer ticker.Stop()

	var lastModTime time.Time
	var lastConfig *gconfig.AppConfiguration // 保存配置数据

	for range ticker.C {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			glog.Log.Errorf("Error reading config file: %v", err)
			continue
		}

		if fileInfo.ModTime() != lastModTime {
			lastModTime = fileInfo.ModTime()

			newConfig, err := LoadAppConfig()
			if err != nil {
				glog.Log.Errorf("Error loading config: %v", err)
				// todo 更新加载App配置数据失败，需报警
				continue
			}

			// 检查新配置与旧配置是否相同，避免不必要的重新加载
			gconfig.AppConfigMutex.Lock()
			if !reflect.DeepEqual(lastConfig, newConfig) {
				lastConfig = newConfig
				// 在这里使用最新的配置数据进行处理
				glog.Log.Errorf("Loaded new config: %+v", newConfig)
				gconfig.AppConfig = newConfig
			}
			gconfig.AppConfigMutex.Unlock()
		}
	}
}
