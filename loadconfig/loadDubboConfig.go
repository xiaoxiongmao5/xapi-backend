package loadconfig

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
	gconfig "xj/xapi-backend/g_config"
	_ "xj/xapi-backend/rpc_api_service"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

func LoadDubboConfig() error {
	// 使用命令行参数来指定配置文件路径
	configFile := flag.String("config", "conf/dubbogo.yaml", "Path to Dubbo-go config file")
	flag.Parse()

	// 设置 DUBBO_GO_CONFIG_PATH 环境变量
	os.Setenv("DUBBO_GO_CONFIG_PATH", *configFile)

	// 加载 Dubbo-go 的配置文件，根据环境变量 DUBBO_GO_CONFIG_PATH 中指定的配置文件路径加载配置信息。配置文件通常包括 Dubbo 服务的注册中心地址、协议、端口等信息。
	if err := config.Load(); err != nil {
		return err
	}
	return nil
}

// 参数: Nacos 服务地址和端口
func RegisterServiceToNacos() {
	nacosHost := gconfig.AppConfig.Nacos.Host
	nacosPort := gconfig.AppConfig.Nacos.Port
	// 最大尝试次数和当前尝试次数
	maxAttempts := 30
	attempt := 1

	// 循环检查 Nacos 服务是否可用
	for attempt <= maxAttempts {
		url := fmt.Sprintf("http://%s:%d/nacos/health", nacosHost, nacosPort)
		resp, err := http.Get(url)
		if err == nil && resp != nil && resp.StatusCode == http.StatusOK {
			fmt.Println("Nacos is up and running, starting backend service...")
			LoadDubboConfig()
			break
		} else {
			fmt.Printf("Attempt %d: Nacos is not ready yet, waiting...\n", attempt)
			attempt++
			time.Sleep(5 * time.Second)
		}
	}

	if attempt > maxAttempts {
		fmt.Println("Max attempts reached. Nacos may not be available.")
		// 在这里可以添加适当的错误处理或退出逻辑
	} else {
		// Nacos 可用后执行启动后端服务的操作
		// 示例：启动后端服务
		// startBackendService()
	}
}
