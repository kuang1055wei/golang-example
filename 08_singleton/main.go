package main

import (
	"crypto/md5"
	rand2 "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
	//rand.
	num := strconv.Itoa(rand.Intn(1000000))
	fmt.Println(num)
	data := []byte(num)
	s := fmt.Sprintf("%x", md5.Sum(data))
	fmt.Println(s)

	// 也可以用这种方式
	h := md5.New()
	h.Write(data)
	s = hex.EncodeToString(h.Sum(nil))
	fmt.Println(s)

	start := time.Now().UnixNano()
	for i := 0; i < 5; i++ {
		x := rand.Intn(100)
		fmt.Println(x)
	}
	fmt.Println(time.Now().UnixNano() - start)
	fmt.Println("---------------")
	start2 := time.Now().UnixNano()
	//使用crypto/rand
	for i := 0; i < 5; i++ {
		x,_ := rand2.Int(rand2.Reader , big.NewInt(100))
		fmt.Println(x)
	}
	fmt.Println(time.Now().UnixNano() - start2)

}
