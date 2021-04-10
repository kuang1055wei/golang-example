package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
1.1.1. runtime.Gosched()
让出CPU时间片，重新等待安排任务(
大概意思就是本来计划的好好的周末出去烧烤，但是你妈让你去相亲,两种情况第一就是你相亲速度非常快，见面就黄不耽误你继续烧烤，
第二种情况就是你相亲速度特别慢，见面就是你侬我侬的，耽误了烧烤，但是还馋就是耽误了烧烤你还得去烧烤)
 */
//func main() {
//	go func(s string) {
//		for i := 0; i < 2; i++ {
//			fmt.Println(s)
//		}
//	}("world")
//	// 主协程
//	for i := 0; i < 2; i++ {
//		// 切一下，再次分配任务
//		runtime.Gosched()
//		fmt.Println("hello")
//	}
//}

/*
1.1.2. runtime.Goexit()
退出当前协程(一边烧烤一边相亲，突然发现相亲对象太丑影响烧烤，果断让她滚蛋，然后也就没有然后了)
 */
//func main() {
//	go func() {
//		defer fmt.Println("A.defer")
//		func() {
//			defer fmt.Println("B.defer")
//			// 结束协程
//			runtime.Goexit()
//			defer fmt.Println("C.defer")
//			fmt.Println("B")
//		}()
//		fmt.Println("A")
//	}()
//	for {
//	}
//}

/*
1.1.3. runtime.GOMAXPROCS
Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。
例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）。

Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
 */

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}