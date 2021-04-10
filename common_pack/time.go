package main

import (
	"fmt"
	"time"
)

func timestampDemo() {
	now := time.Now()           //获取当前时间
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("当前时间:%v\n" , now)
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	fmt.Println("MM-DD-YYYY : ", now.Format("01-02-2006"))
	fmt.Println(now.Format("2006.01.02 15:04:05"))
}
func main() {
	timestampDemo()
}
