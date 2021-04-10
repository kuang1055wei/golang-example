package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int
	Weight float32
}

func main() {
	p := Person{
		Name:   "kw",
		Age:    26,
		Weight: 55.5,
	}
	json1 , _ := json.Marshal(p)
	fmt.Println(string(json1))
	fmt.Println(fmt.Sprintf("%s" , json1))

	//json转struct
	var p2 Person
	_ = json.Unmarshal(json1 , &p2)
	fmt.Printf("%+v\n" , p2)

	//json转map
	var p3 map[string]interface{}
	_ = json.Unmarshal(json1 , &p3)
	fmt.Printf("%v\n" , p3)
}
