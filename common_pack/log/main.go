package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	//log.Fatalln("这是一条会触发fatal的日志。")

	var buf [16]byte
	os.Stdin.Read((buf[:]))
	f,err := os.Create("a")
	fmt.Println(f , err)
}
