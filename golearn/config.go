package main

import (
	"fmt"
	"os"
)

type Config struct {
	Host     string
	User     string
	DataPath string
}

func loadConfig() Config {
	return Config{
		Host:     os.Getenv("JUMPSERVER_HOST"),
		User:     os.Getenv("SSH_USER"),
		DataPath: os.Getenv("DATA_PATH"),
	}
}

func main() {
	cfg := loadConfig()

	if cfg.DataPath == "" {
		fmt.Println("错误：没有配置 DATA_PATH，程序退出")
		return
	}

	fmt.Println("配置加载成功：")
	fmt.Println("  跳板机:", cfg.Host)
	fmt.Println("  用户:", cfg.User)
	fmt.Println("  数据路径:", cfg.DataPath)
}
