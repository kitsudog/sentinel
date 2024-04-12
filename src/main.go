package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(`
Usage: 
	sentinel <port> <cmd> [params...]
`))
		os.Exit(1)
	}
	port := os.Args[1]
	ln, err := net.Listen("tcp", ":"+port) // 请将12345换成你想监听的端口
	if err != nil {
		fmt.Printf("listen fail: %v\n", err)
		return
	}
	defer ln.Close()

	fmt.Printf("listen %v...\n", port)

	_, err = ln.Accept()
	if err != nil {
		fmt.Printf("accept fail: %v\n", err)
		return
	}

	command := os.Args[2]
	args := os.Args[3:]

	cmd := exec.Command(command, args...)
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting command: %s\n", err)
		return
	} else {
		fmt.Printf("hello world")
		return
	}
}
