package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//处理错误函数
func checkErrs(err error) {
	if err != nil {
		panic(err)
	}
}

func brands() {
	//读取Api数据
	resp, err := http.Get("https://j.autohome.com.cn/pcplatform/common/getBrandInfo")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	//json转化成map
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(data), &m)
	//fmt.Println(m["result"])
	fmt.Printf("%T\n" , m);
	fmt.Printf("%T\n" , m["result"].(map[string]interface{})["A"]);
	//fmt.Println(m["result"].(map[string]interface{})["A"].([]interface{})[3].(map[string]interface{})["name"])
}

func main() {
	brands()
}