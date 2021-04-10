package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int64
	Weight float64
}

func NewPerson(name string , age int64,weight float64) (*Person)  {
	return &Person{
		Name: name,
		Age:age,
		Weight: weight,
	}
}

func main() {
	//p1 := Person{
	//	Name:   "七米",
	//	Age:    18,
	//	Weight: 71.5,
	//}
	p1 := NewPerson("七米", 18 , 71.5)
	// struct -> json string
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// json string -> struct
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
