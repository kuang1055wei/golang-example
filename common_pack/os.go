package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var rwLock sync.RWMutex

func main() {

	//f, _ := os.OpenFile("./hash.go" , os.O_RDWR , 0777)
	//f, _ := os.Open("./a.txt")
	//defer f.Close()
	//readall , _ := io.ReadAll(f)
	//fmt.Println("读取所有" , string(readall))

	//fmt.Println("read读取")
	//buf := make([]byte , 1)
	//for  {
	//	n,_ := f.Read(buf)
	//	fmt.Println(string(buf), "读取子节:", n)
	//	if n == 0 {
	//		break
	//	}
	//}

	//fmt.Println("read读取2")
	//buf2 := bufio.NewReader(f)
	//for  {
	//	b, _, c := buf2.ReadLine()
	//	if c==io.EOF {
	//		break
	//	}
	//	fmt.Println(string(b))
	//}

	//io写入
	//f3,_ := os.OpenFile("./a.txt" , os.O_APPEND , 0666)
	//n,_ := io.WriteString(f3 , "\nbuffer")
	//fmt.Printf("写入 %d 个字节n\n", n)

	//bufio写入
	//f2 , _ := os.OpenFile("./a.txt" , os.O_APPEND , 0777)
	//w := bufio.NewWriter(f2)
	//n3, _ := w.WriteString("buffer")
	//fmt.Printf("写入 %d 个字节n", n3)
	//_ = w.Flush()//写入文件
	//err := f2.Sync()//同步磁盘
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer f2.Close()

	//批量写入
	f3 , _ := os.OpenFile("./a.txt" , os.O_APPEND , os.ModePerm)
	defer f3.Close()
	//w1 := bufio.NewWriter(f3)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			str := "\nwrite_"+strconv.Itoa(i)
			n,_ := f3.WriteString(str)
			fmt.Printf("写入 %d 个字节n", n)
			wg.Done()
		}(i)
	}
	//w1.Flush()
	//f3.Sync()
	wg.Wait()

}
