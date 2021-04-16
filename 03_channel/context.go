package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

//停止协程和context的使用
func main() {
	//协程使用案例
	go MyChan()

	//go StopChan()
	//
	//go contextStop()
	time.Sleep(10 * time.Second)
}

func MyChan() {
	id1, _ := strconv.Atoi("3")
	id2, _ := strconv.Atoi("4")
	art1Chan := getArt(id1)
	art2Chan := getArt(id2)

	//art1 := <-art1Chan
	//art2 := <-art2Chan
	data := make(map[string]interface{})
	data["art1"] = <-art1Chan
	data["art2"] = <-art2Chan
	fmt.Printf("%+v\n", data)
}

func getArt(id int) <-chan string {
	artChan := make(chan string)
	go func(id int) {
		art := "article_" + strconv.Itoa(id)
		artChan <- art
	}(id)
	return artChan
}

//协程退出
func StopChan() {
	var wg sync.WaitGroup
	wg.Add(1)
	stopCh := make(chan bool) //停止监控的
	go func() {
		defer wg.Done()
		watchDog(stopCh, "【监控狗】")
	}()
	time.Sleep(5 * time.Second)
	stopCh <- true //发送停止指令
	wg.Wait()
}

//使用通知指令停止协程
func watchDog(stopCh chan bool, name string) {
	for {
		select {
		case <-stopCh:
			fmt.Println(name, "停止指令收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控")
		}
		time.Sleep(1 * time.Second)
	}
}

func contextStop() {
	//WithCancel (parent Context)：生成一个可取消的 Context。
	//WithDeadline (parent Context, d time.Time)：生成一个可定时取消的 Context，参数 d 为定时取消的具体时间。
	//WithTimeout (parent Context, timeout time.Duration)：生成一个可超时取消的 Context，参数 timeout 用于设置多久后取消
	//WithValue (parent Context, key, val interface {})：生成一个可携带 key-value 键值对的 Context。
	//取消多个 如下
	//wg.Add(3)
	//ctx,stop:=context.WithCancel(context.Background())
	//go func() {
	//	defer wg.Done()
	//	watchDog(ctx,"【监控狗2】")
	//}()
	//go func() {
	//	defer wg.Done()
	//	watchDog(ctx,"【监控狗3】")
	//}()
	//stop()
	//取消多个
	var wg sync.WaitGroup
	wg.Add(2)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog2(ctx, "【监控狗context】")
	}()
	time.Sleep(5 * time.Second)
	stop() //停止监控狗
	wg.Wait()
}

//使用context停止协程
func watchDog2(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控")
		}
		time.Sleep(1 * time.Second)
	}
}
