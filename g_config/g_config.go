package gconfig

import (
	"sync"
)

// 用于给用户分配accessKey,secretKey
const SALT = "xj"

// 用于给用户生成登录验证token（jwt）
const SecretKey = "your-secret-key"

// App配置数据
type AppConfiguration struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Dbname   string `json:"dbname"`
		Username string `json:"username"`
		Password string `json:"password"`
		SavePath string `json:"savePath"`
	} `json:"database"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
	IPWhiteList []string `json:"ipWhiteList"`
	IPBlackList []string `json:"ipBlackList"`
	Nacos       struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"nacos"`
}

var (
	AppConfigMutex sync.Mutex
	AppConfig      *AppConfiguration
)
