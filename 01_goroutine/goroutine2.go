package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string)  {
	ch <- "1"
	ch <- "2"
	ch <- "3"
	ch <- "4"
	ch <- "5"
	ch <- "6"
	close(ch)
}

func getData(ch chan string)  {
	//for  {
	//	input,ok := <- ch
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(input)
	//}
	//第二种方式
	for input := range ch {
		fmt.Println(input)
	}
}
