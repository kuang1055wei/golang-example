package mysqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"strconv"
	"time"
)

//type Article struct {
//	Id           int64     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
//	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
//	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
//	DeletedAt    time.Time `gorm:"column:deleted_at" json:"deleted_at"`
//	Title        string    `gorm:"column:title;NOT NULL" json:"title"`
//	Cid          int64     `gorm:"column:cid;NOT NULL" json:"cid"`
//	Desc         string    `gorm:"column:desc" json:"desc"`
//	Content      string    `gorm:"column:content" json:"content"`
//	Img          string    `gorm:"column:img" json:"img"`
//	CommentCount int64     `gorm:"column:comment_count;default:0;NOT NULL" json:"comment_count"`
//	ReadCount    int64     `gorm:"column:read_count;default:0;NOT NULL" json:"read_count"`
//}
type Article struct {
	Id           int64      `db:"id"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
	Title        string     `db:"title"`
	Cid          int64      `db:"cid"`
	Desc         *string    `db:"desc"    json:",omitempty"`
	Content      *string    `db:"content"    json:",omitempty"`
	Img          *string    `db:"img"    json:",omitempty"`
	CommentCount int64      `db:"comment_count"`
	ReadCount    int64      `db:"read_count"`
}

func StructQueryFiled(id int64) Article {
	var article Article
	_ = MysqlDb.Select(&article, "select id,title from article where id=?", id)
	return article
}

func StructQueryAllField(num int64) []Article {
	articles := make([]Article, 0)
	var columns string
	columns = "id,title,content,read_count"
	////columns = "*"
	//if err := MysqlDb.Select(&articles, "SELECT "+columns+" FROM article where id >3  limit ?", num); err != nil {
	//	fmt.Println(err)
	//	return articles
	//}

	//使用name
	//stmt , _ := MysqlDb.PrepareNamed(`select `+columns+` from article where id>:id limit :num`)
	//arg := map[string]interface{}{
	//	"id":3,
	//	"num":num,
	//}
	//for key, val := range arg {
	//	fmt.Println(key,"-----------",val)
	//}
	//err := stmt.Select(&articles, arg)
	//if err!=nil {
	//	fmt.Println(err)
	//	return articles
	//}

	//arg := map[string]interface{}{
	//	"id" : []int{3,4,5},
	//	//"num" : num,
	//}
	//query, args, _ := sqlx.Named("select "+columns+" from article where id in (:id) limit :num", arg)
	query, args, _ := sqlx.In("select "+columns+" from article where id in (?) limit ?", []int{3,4,5} , 10)
	fmt.Println(args)
	fmt.Println(query)
	query = MysqlDb.Rebind(query)
	err := MysqlDb.Select(&articles, query , args...)
	if err!=nil {
		fmt.Println(err)
		return articles
	}
	return articles
}

func StructInsert() int64 {
	ret, err := MysqlDb.Exec("insert into article(title,content,cid) values (?,?,?)", "测试添加", "测试添加内容。。。", 0)
	if err != nil {
		panic(err)
	}
	id, _ := ret.LastInsertId()
	return id
}

// 更新数据
func StructUpdate(id int64, title string) int64 {

	ret, _ := MysqlDb.Exec("UPDATE article set title=? where id=?", title, id)
	upd_nums, _ := ret.RowsAffected()

	return upd_nums
}

// 删除数据
func StructDel(id int64) bool {

	_, err := MysqlDb.Exec("delete from article where id=?", id)
	//_,err := ret.RowsAffected()
	if err != nil {
		return false
	}
	return true
}

// 查询数据,取所有字段,不采用结构体
func RawQueryAllField() {

	//查询数据，取所有字段
	rows2, _ := MysqlDb.Query("select * from article")

	//返回所有列
	cols, _ := rows2.Columns()

	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))

	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))
	//这里scans引用vals，把数据填充到[]byte里
	for k := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...)
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	fmt.Println(result)
}

func Trans()  {
	tx , err := MysqlDb.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	//rand.Seed(time.Now().UnixNano())
	//randInt := rand.Intn(10000)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randInt := r.Intn(10000)
	_ , err = MysqlDb.Exec("UPDATE article set title=? where id=?", "修改标题" +strconv.Itoa(randInt), 5)
	if err != nil {
		fmt.Println(err)
		_ = tx.Rollback()
	}

	_ = tx.Commit()
}
