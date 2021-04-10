package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"go_learn/gplearn.com/mysql/myCrud"
	"go_learn/gplearn.com/mysql/mysqlx"
)



func main() {
	//art := mysqlx.StructQueryFiled(1)
	////fmt.Println(art.Id , art.Title)
	//j2,_ := json.Marshal(art)
	//fmt.Println(string(j2))

	arts := mysqlx.StructQueryAllField(100)
	fmt.Println(len(arts))
	json3 , _ := json.Marshal(arts)
	fmt.Println(string(json3))


	mysqlx.Trans()

	//art := myCrud.StructQueryFiled(1)
	////fmt.Println(art.Id , art.Title)
	//j2,_ := json.Marshal(art)
	//fmt.Println(string(j2))

	//artList := myCrud.StructQueryAllField(100)
	//j, _ := json.Marshal(artList)
	////jStr := fmt.Sprintf("%#v" , j)
	//fmt.Println(string(j))
	//for _, article := range artList {
	//	fmt.Println(article.Id, article.Title , article.ReadCount)
	//}


	//myCrud.StructInsert()

	//num := myCrud.StructUpdate(4 , "修改后的标题")
	//fmt.Println("修改成功，影响行数:",num)
	//
	//delRes := myCrud.StructDel(6)
	//fmt.Println("删除结果为:",delRes)
	//
	//myCrud.RawQueryAllField()
}

