package main

import (
	"fmt"
	"os"
)

type Server struct {
	Name   string
	IP     string
	Port   int
	Online bool
}

func main() {

	jumpHost := os.Getenv("JUMP_HOST")
	if jumpHost == "" {
		jumpHost = "10.0.0.1"
	}
	sshUser := os.Getenv("SSH_USER")
	if sshUser == "" {
		sshUser = "zzy"
	}
	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		dataPath = "/data/2026-10-07"
	}
	fmt.Println("JUMP_HOST:", jumpHost)
	fmt.Println("SSH_USER:", sshUser)
	fmt.Println("DATA_PATH:", dataPath)

	servers := []Server{
		{
			Name:   "Server1",
			IP:     "10.0.0.1",
			Port:   22,
			Online: true,
		},
		{
			Name:   "Server2",
			IP:     "10.0.0.2",
			Port:   22,
			Online: false,
		},
		{
			Name:   "Server3",
			IP:     "10.0.0.3",
			Port:   22,
			Online: false,
		},
	}
	for i, server := range servers {
		if server.Online {
			fmt.Printf("Name %d: %s\n 从 IP: %s 拉取 dataPath: %s\n", i+1, server.Name, server.IP, dataPath)
		}
		if !server.Online {
			fmt.Printf("Name %d: %s 跳过，不在线\n", i+1, server.Name)
		}
	}
}
