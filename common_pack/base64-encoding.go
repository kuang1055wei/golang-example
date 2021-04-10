package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "http://www.baidu.com"

	// Go 同时支持标准的和 URL 兼容的 base64 格式。
	// 编码需要使用 `[]byte` 类型的参数，所以要将字符串转成此类型。
	str := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(str)


	// 解码可能会返回错误，如果不确定输入信息格式是否正确，
	// 那么，你就需要进行错误检查了。
	str2, _ := base64.StdEncoding.DecodeString(str)
	fmt.Println(string(str2))

	// 使用 URL 兼容的 base64 格式进行编解码。
	urlStr := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlStr)
	urlStr2,_ := base64.URLEncoding.DecodeString(urlStr)
	fmt.Println(string(urlStr2))

}
