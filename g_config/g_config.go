package gconfig

// 用于给用户分配accessKey,secretKey
const SALT = "xj"

// 用于给用户生成登录验证token（jwt）
const SecretKey = "your-secret-key"

type AppConfiguration struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Dbname   string `json:"dbname"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Server struct {
		Port int `json:"port"`
	} `json:"server"`
}

var AppConfig AppConfiguration
