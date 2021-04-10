package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"gorm.io/plugin/dbresolver"
)


type Result struct {
	ID   int
	Title string
	Content  string
}


func main() {
	dsn := "root:root@tcp(192.168.50.100)/ginblog?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDb := mysql.Open(dsn)
	db, err := gorm.Open(sqlDb, &gorm.Config{})
	//sqlDB,err := db.DB()
	//sqlDB.SetMaxIdleConns(10)
	//连接池#
	db.Use(
		dbresolver.Register(dbresolver.Config{ /* xxx */ }).
		SetConnMaxIdleTime(time.Hour).
		SetConnMaxLifetime(24 * time.Hour).
		SetMaxIdleConns(100).SetMaxOpenConns(200),
	)

	if err != nil {

		panic("failed to connect database")
	}

	var result Result
	db.Raw("select id,title from article where id=?" , 1).Scan(&result)

	fmt.Println(result)

	result2 := map[string]interface{}{}
	db.Table("article").Select("id , title").Take(&result2)
	fmt.Println(result2)

	db.Find(&result , []int{1,2,3})
	fmt.Println(result2)
	//// Migrate the schema
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}