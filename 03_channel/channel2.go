package main

import (
	"fmt"
	"strconv"
	"time"
)

//var wg sync.WaitGroup

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string)  {
	for i := 0; i < 10; i++ {
		ch <- strconv.Itoa(i)
	}
}

func getData(ch chan string)  {
	var input string
	for{
		input = <- ch
		fmt.Println(input)
	}
}
