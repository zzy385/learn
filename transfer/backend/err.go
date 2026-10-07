package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("/workspaces/learn/transfer/fake_remote/data/2026-10-07/result.txt")
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}

	fmt.Println("读取文件成功:", string(data))
}
