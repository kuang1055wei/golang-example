package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "an;jlkjlkajskd"
	md5Str := md5.Sum([]byte(str))
	fmt.Println(md5Str)
	fmt.Println(fmt.Sprintf("%x" , md5Str))

	//方式二
	str2 := "lahslkdhajshkjlahs"
	h := md5.New()
	h.Write([]byte(str2))
	md5str2 := hex.EncodeToString(h.Sum(nil))
	fmt.Println(md5str2)
}
