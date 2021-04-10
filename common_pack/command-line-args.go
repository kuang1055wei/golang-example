package main

import (
	"fmt"
	"os"
)

//1:go build command-line-args.go
//2:os_args a b c d e
func main() {
	// `os.Args` 提供原始命令行参数访问功能。
	// 注意，切片中的第一个参数是该程序的路径，
	// 并且 `os.Args[1:]` 保存程序的所有参数。
	args := os.Args
	argsParam := args[1:]

	// 你可以使用标准的索引位置方式取得单个参数的值。
	arg := os.Args[3]

	fmt.Println(args)
	fmt.Println(argsParam)
	fmt.Println(arg)
}
