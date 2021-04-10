package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"sync"
	"time"
)
var once sync.Once


func md51()  {
	num := strconv.Itoa(rand.Intn(1000000))
	data := []byte(num)
	s := fmt.Sprintf("%x", md5.Sum(data))
	fmt.Println(s)
}

func md52()  {
	num := strconv.Itoa(rand.Intn(1000000))
	data := []byte(num)
	// 也可以用这种方式
	h := md5.New()
	h.Write(data)
	s := hex.EncodeToString(h.Sum(nil))
	fmt.Println(s)
}

func md53()  {
	cruTime := time.Now().Unix()
	h := md5.New()
	_ , _ = io.WriteString(h, strconv.FormatInt(cruTime, 10))
	token := fmt.Sprintf("%x" , h.Sum(nil))
	fmt.Println(token)
}

func main() {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
	md51()
	md52()
	md53()
}
