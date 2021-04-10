package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	name string
	age  uint
}

type xx struct {
	Person
	aa []string
	b  int
}

func test() {
	fmt.Println("hello word")
}

func main() {
	var name string = "哈哈哈哈"
	go test()
	fmt.Println("我是主线程")
	go func(name string) {
		fmt.Println(name)
		fmt.Println(rand.Intn(100))
	}(name)
	time.Sleep(time.Second)

}
