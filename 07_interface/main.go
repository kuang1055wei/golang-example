package main

import "fmt"

//func main() {
//	//空接口
//	var studentInfo = make(map[string]interface{})
//	studentInfo["name"] = "kw"
//	studentInfo["age"] = 10
//	studentInfo["sex"] = "男"
//	studentInfo["married"] = false
//	fmt.Println(studentInfo)
//
//	show("anc")
//	show(map[string]string{
//		"1":"哈哈哈哈",
//		"2":"呵呵呵呵",
//	})
//	show([]int{1,2,3})
//	show([...]string{"1","2","3","4"})
//	show(int64(123456789))
//	show(func() {
//
//	})
//}

//断言
func main() {
	var x interface{}
	x = "Hello 沙河"
	v,ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
	a := true
	justifyType(a)

}

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}