package main

import "fmt"

// 定义结构体：一台服务器（这就是一张"登记表"）
type Server struct {
	Name    string // 服务器名
	Host    string // IP 地址
	Port    int    // 端口
	HasData bool   // 有没有数据
}

func main() {
	// 方式一：按字段顺序填
	s1 := Server{"超算A100", "10.0.0.1", 22, true}

	// 方式二：指名道姓地填（推荐，顺序错了也没事）
	s2 := Server{Name: "下级服务器1", Host: "192.168.1.10", Port: 22}

	// 用"点号 ."取字段
	fmt.Println("服务器1：", s1.Name, s1.Host, s1.Port, s1.HasData)
	fmt.Println("服务器2：", s2.Name, s2.Host, s2.Port, s2.HasData)
}
