package main

import (
	"log"
	"net"
	"fmt"
	// "bufio"
)

func main() {
	port := ":8080"
	
	// localhost の8080ポートを解放
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("can not listen", port)
	}
	fmt.Println("listening ", "port")

	// connectionの受付
	conn, err := listen.Accept()
	if err != nil {
		log.Fatal("can not established connection")
	}
	// メモリ割り当て byte型はuint8型のエイリアス
	buf := make([]byte, 1024)

	fmt.Printf("[Remote Address]\n%s\n", conn.RemoteAddr())
	n, err := conn.Read(buf) // nはリクエストのバイト列のサイズ
	if err != nil {
		log.Fatal("can not read request")
	}
	fmt.Printf("[Message]\n%s", string(buf[:n]))

	conn.Write([]byte("connected!"))

	conn.Close()
}