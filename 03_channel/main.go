package main

import "fmt"

/**
单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存进行数据交换，但是共享内存在不同的goroutine中容易发生竞态问题。为了保证数据交换的正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题。

Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。

Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
 */

/*
1.1.2. channel类型
channel是一种类型，一种引用类型。声明通道类型的格式如下：

    var 变量 chan 元素类型

举几个例子：

    var ch1 chan int   // 声明一个传递整型的通道
    var ch2 chan bool  // 声明一个传递布尔型的通道
    var ch3 chan []int // 声明一个传递int切片的通道
 */

//func main() {
//	var ch chan int
//	fmt.Println(ch) // nil
//	//声明的通道后需要使用make函数初始化之后才能使用。
//	//make(chan 元素类型, [缓冲大小]),channel的缓冲大小是可选的。
//	ch4 := make(chan int)
//	ch5 := make(chan bool)
//	ch6 := make(chan []int)
//	fmt.Println(ch4 , ch5 , ch6) // nil
//}

/*
1.1.4. channel操作
通道有发送（send）、接收(receive）和关闭（close）三种操作。

发送和接收都使用<-符号。

现在我们先使用以下语句定义一个通道：
 */
//func main() {
//	ch := make(chan int)
//	ch <- 10 // 把10发送到ch中
//
//	x := <- ch // 从ch中接收值并赋值给变量x
//	//<-ch       // 从ch中接收值，忽略结果
//	fmt.Println(x)
//	//   1.对一个关闭的通道再发送值就会导致panic。
//	//    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
//	//    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//	//    4.关闭一个已经关闭的通道会导致panic。
//	//defer close(ch)
//
//}

//为什么会出现deadlock错误呢？
//因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。
//就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送。
//上面的代码会阻塞在ch <- 10这一行代码形成死锁
//func main() {
//	ch := make(chan int)
//	ch <- 10
//	fmt.Println("发送成功")
//}

//解决上面的问题
//func recv(c chan int) {
//	ret := <-c
//	fmt.Println("接收成功", ret)
//}
//func main() {
//	ch := make(chan int)
//	go recv(ch) // 启用goroutine从通道接收值
//	ch <- 10
//	fmt.Println("发送成功")
//}


//有缓冲的通道
//只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。
//就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。
//
//我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。
//func main() {
//	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
//	ch <- 10
//	fmt.Println("发送成功")
//}

//1.1.7. close()
//可以通过内置的close()函数关闭channel（如果你的管道不往里存值或者取值的时候一定记得关闭管道）
//func main() {
//	c := make(chan int)
//	go func() {
//		for i := 0; i < 5; i++ {
//			c <- i
//		}
//		close(c)
//	}()
//	for {
//		if data, ok := <-c; ok {
//			fmt.Println(data)
//		} else {
//			break
//		}
//	}
//	fmt.Println("main结束")
//}

// 判断一个通道是否被关闭了
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	// 开启goroutine将0~100的数发送到ch1中
//	go func() {
//		for i := 0; i < 100; i++ {
//			ch1 <- i
//		}
//		close(ch1)
//	}()
//	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
//	go func() {
//		for {
//			i, ok := <-ch1 // 通道关闭后再取值ok=false
//			if !ok {
//				break
//			}
//			ch2 <- i * i
//		}
//		close(ch2)
//	}()
//	// 在主goroutine中从ch2中接收值打印
//	for i := range ch2 { // 通道关闭后会退出for range循环
//		fmt.Println(i)
//	}
//}

/*
1.1.9. 单向通道
有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。
 */
// 1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
// 2.<-chan int是一个只能接收的通道，可以接收但是不能发送。
//在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的。
//只能发送
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}
//只能发送
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
//只能接收
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}