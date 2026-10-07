package main

import "fmt"

type Server struct {
	Name   string
	IP     string
	Port   int
	Online bool
}

func main() {
	servers := []Server{
		{
			Name:   "Server1",
			IP:     "192.168.1.1",
			Port:   22,
			Online: true,
		},
		{
			Name:   "Server2",
			IP:     "192.168.1.2",
			Port:   23,
			Online: false,
		},
		{
			Name:   "Server3",
			IP:     "192.168.1.3",
			Port:   24,
			Online: false,
		},
	}
	for i, server := range servers {
		fmt.Printf("Server %d:\n", i+1)
		fmt.Println(server)
	}
}
