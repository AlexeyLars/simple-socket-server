package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err == nil {
		fmt.Println("Connected to localhost:8080")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message == "OK" {
			fmt.Println("OK")
		}
	}
}
