package main

import (
	"fmt"
	"strconv"
	"sync"
)

//1.1.1. sync.WaitGroup
//在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：
//
//方法名	功能
//(wg * WaitGroup) Add(delta int)	计数器+delta
//(wg *WaitGroup) Done()	计数器-1
//(wg *WaitGroup) Wait()	阻塞直到计数器变为0
//sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。
//
//我们利用sync.WaitGroup将上面的代码优化一下：
//var wg sync.WaitGroup
//
//func hello() {
//	defer wg.Done()
//	fmt.Println("Hello Goroutine!")
//}
//func main() {
//	wg.Add(1)
//	go hello() // 启动另外一个goroutine去执行hello函数
//	fmt.Println("main goroutine done!")
//	wg.Wait()
//}

//sync.Once
//说在前面的话：这是一个进阶知识点。
//
//在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次，例如只加载一次配置文件、只关闭一次通道等。
//
//Go语言中的sync包中提供了一个针对只执行一次场景的解决方案–sync.Once。
//
//sync.Once只有一个Do方法，其签名如下：
//
//func (o *Once) Do(f func()) {}
//注意：如果要执行的函数f需要传递参数就需要搭配闭包来使用。
//加载配置文件示例
//延迟一个开销很大的初始化操作到真正用到它的时候再执行是一个很好的实践。因为预先初始化一个变量（比如在init函数中完成初始化）
//会增加程序的启动耗时，而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。我们来看一个例子：

//var icons map[string]image.Image
//var loadIconsOnce sync.Once
//
//func loadIcons() {
//	icons = map[string]image.Image{
//		"left":  loadIcon("left.png"),
//		"up":    loadIcon("up.png"),
//		"right": loadIcon("right.png"),
//		"down":  loadIcon("down.png"),
//	}
//}
//
//// Icon 是并发安全的
//func Icon(name string) image.Image {
//	loadIconsOnce.Do(loadIcons)
//	return icons[name]
//}


//sync.Map
//Go语言中内置的map不是并发安全的
//var m = make(map[string]int)
//
//func get(key string) int {
//	return m[key]
//}
//
//func set(key string, value int) {
//	m[key] = value
//}
//
//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			set(key, n)
//			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}
//上面的代码开启少量几个goroutine的时候可能没什么问题，当并发多了之后执行上面的代码就会报fatal error: concurrent map writes错误。

//像这种场景下就需要为map加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。
//开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}