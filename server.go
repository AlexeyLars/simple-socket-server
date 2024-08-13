package main

import (
	"fmt"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close() // закрываем сокет по окончании обработки
	defer fmt.Println("Connection with", conn.RemoteAddr(), "closed")

	conn.Write([]byte("OK"))
	return
}

func main() {
	fmt.Println("Start server...")
	// открываем слушающий сокет
	listener, _ := net.Listen("tcp", "localhost:8080")

	for {
		conn, err := listener.Accept() // принимаем TCP-соединение с клиентом и создаём новый сокет
		if err != nil {
			continue
		}
		fmt.Println("New connection from", conn.RemoteAddr())
		// если не возникло никаких ошибок, то обрабатываем соединение
		go handler(conn)
	}
}
