package myCrud

import (
	"fmt"
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
	Id           int64     `db:"id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	DeletedAt    time.Time `db:"deleted_at"`
	Title        string    `db:"title"`
	Cid          int64     `db:"cid"`
	Desc         string    `db:"desc"`
	Content      string    `db:"content"`
	Img          string    `db:"img"`
	CommentCount int64     `db:"comment_count"`
	ReadCount    int64     `db:"read_count"`
}

func StructQueryFiled(id int64) (Article)  {
	row := MysqlDb.QueryRow("select id,Title from article where id=?" , id)
	article := new(Article)
	if err := row.Scan(&article.Id,&article.Title); err!=nil{
		fmt.Print("查询失败:%v\n" , err)
		return Article{}
	}
	return *article
}

func StructQueryAllField(num int64) []Article  {
	articles := make([]Article , 0)
	var columns string = "id,title,content,read_count"
	rows , _ := MysqlDb.Query("SELECT "+columns+" FROM article limit ?" , num)
	var art Article
	for rows.Next(){
		_ = rows.Scan(&art.Id, &art.Title, &art.Content,&art.ReadCount)
		articles = append(articles , art)
	}
	return articles
}



func StructInsert() int64 {
	ret,err:= MysqlDb.Exec("insert into article(title,content,cid) values (?,?,?)" , "测试添加","测试添加内容。。。",0)
	if err != nil {
		panic(err)
	}
	id , _ := ret.LastInsertId()
	return id
}

// 更新数据
func StructUpdate(id int64 , title string) int64 {

	ret,_ := MysqlDb.Exec("UPDATE article set title=? where id=?",title,id)
	upd_nums,_ := ret.RowsAffected()

	return upd_nums
}

// 删除数据
func StructDel(id int64) bool {

	_ , err := MysqlDb.Exec("delete from article where id=?",id)
	//_,err := ret.RowsAffected()
	if err!=nil {
		return false
	}
	return true
}

// 查询数据,取所有字段,不采用结构体
func RawQueryAllField() {

	//查询数据，取所有字段
	rows2, _ := MysqlDb.Query("select * from article");

	//返回所有列
	cols, _ := rows2.Columns();

	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols));

	//这里表示一行填充数据
	scans := make([]interface{}, len(cols));
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k];
	}

	i := 0;
	result := make(map[int]map[string]string);
	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...);
		//每行数据
		row := make(map[string]string);
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k];
			//这里把[]byte数据转成string
			row[key] = string(v);
		}
		//放入结果集
		result[i] = row;
		i++;
	}
	fmt.Println(result);
}