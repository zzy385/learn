package main

import "fmt"

func main() {
	// 创建切片：一列服务器 IP（注意是 []string 不是 string）
	servers := []string{"192.168.1.10", "192.168.1.11", "192.168.1.12"}

	// 追加一台（新服务器上线了）——append 要接住返回值
	servers = append(servers, "192.168.1.13")

	fmt.Println("一共", len(servers), "台服务器")

	// 遍历方式一：for + 下标（i 从 0 数到 len-1）
	for i := 0; i < len(servers); i++ {
		fmt.Println("第", i, "台：", servers[i])
	}

	// 遍历方式二：for range（项目里最常用！）
	for i, ip := range servers {
		fmt.Println(i, "→", ip)
	}
}
