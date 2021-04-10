package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter , r *http.Request)  {
	_ = r.ParseForm()  // 解析参数，默认是不会解析的
	fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // 这个写入到 w 的是输出到客户端的
}
func main() {
	http.HandleFunc("/" , sayHello)
	err := http.ListenAndServe(":9000" , nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
}
