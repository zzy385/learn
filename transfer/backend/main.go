package main

import (
	"fmt"
	"os"
)

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
}
