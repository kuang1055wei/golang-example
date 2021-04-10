package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	randMath "math/rand"
	"time"
)

func main() {
	//真随机 crypto/rand
	rand1,_ := rand.Int(rand.Reader , big.NewInt(1000))
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
}
