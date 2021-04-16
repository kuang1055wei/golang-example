package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	bigint, _ := rand.Int(rand.Reader, big.NewInt(100000000))
	fmt.Println(bigint)
}
