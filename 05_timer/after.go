package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//	ch := make(chan error, 1)
//	go func() { ch <- client.Call("Service.Method", args, &reply) }()
//	select {
//	case resp := <-ch:
//		fmt.Println(resp)
//		// use resp and reply
//	case <-time.After(time.Second * 5):
//		// call timed out
//		break
//	}
//
//}
