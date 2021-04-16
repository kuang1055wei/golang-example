package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	randMath "math/rand"
	"strings"
	"time"
)

func main() {
	//真随机 crypto/rand
	rand1, _ := rand.Int(rand.Reader, big.NewInt(1000))
	fmt.Println(rand1)

	//假随机 math/rand
	source := randMath.NewSource(time.Now().UnixNano())
	r := randMath.New(source)
	rand2 := r.Intn(1000)
	fmt.Println(rand2)

	//假随机2用seed产生随机数种子,使用全局的
	randMath.Seed(time.Now().UnixNano())
	rand3 := randMath.Intn(1000)
	fmt.Println(rand3)

	str := randString(5)
	fmt.Println(str, strings.ToLower(str))
}

//随机字符串
func randString(len int) string {
	//ASCII 65-90 =A-Z  97-122=a-z
	r := randMath.New(randMath.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
