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
	fmt.Println("listening ", port)

	for {
		// 接続待ち
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("can not established connection")
		}
		// メモリ割り当て byte型はuint8型のエイリアス
		buf := make([]byte, 1024)

		fmt.Printf("[Remote Address]\n%s\n", conn.RemoteAddr()) // クライアントのアドレスを出力 ex.) 172.17.0.1:54696
		// buf にリクエストを代入する.
		n, err := conn.Read(buf) // nはリクエストのバイト列のサイズらしい
		if err != nil {
			log.Fatal("can not read request")
		}
		fmt.Println("n is ", n)
		fmt.Printf("[Message]\n%s", string(buf[:n]))

		conn.Write([]byte("connected!\n"))
	}
	// conn.Close()
}