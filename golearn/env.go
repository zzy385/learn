package main

import (
	"fmt"
	"os"
)

var JUMPSERVER_HOST = "10.0.0.1"
var SSH_USER = "zzy"
var DATA_PATH = "/data/2026-10-01"

func main() {
	// 去环境变量里查三个配置
	host := os.Getenv("JUMPSERVER_HOST")
	user := os.Getenv("SSH_USER")
	dataPath := os.Getenv("DATA_PATH")

	fmt.Println("跳板机地址:", host)
	fmt.Println("SSH用户:", user)
	fmt.Println("数据路径:", dataPath)
}
