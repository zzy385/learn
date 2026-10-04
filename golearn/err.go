package main

import (
	"fmt"
	"os"
)

func main() {
	// os.ReadFile 读文件：返回两个值（内容, 错误）
	data, err := os.ReadFile("main.go") // 读当前目录下的 main.go 文件

	// 检查错误：err 不是空 = 出错了
	if err != nil {
		fmt.Println("出错了：", err)
		fmt.Println("程序没崩——这就是错误处理的功劳")
		return // 提前结束 main，不再往下走
	}

	fmt.Println("文件内容是：", string(data))
}
