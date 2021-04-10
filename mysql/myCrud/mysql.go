package myCrud

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME = "root"
	PASS_WORD = "root"
	HOST      = "192.168.50.100"
	PORT      = "3306"
	DATABASE  = "ginblog"
	CHARSET   = "utf8"
)

func init() {
	//"root:root@tcp(192.168.50.100:3306)/ginblog?charset=utf8"
	dbDns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDns)
	if MysqlDbErr != nil {
		log.Print("dbDbs:" + dbDns)
		panic("数据源配置不正确:" + MysqlDbErr.Error())
	}
	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
}

