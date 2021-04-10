package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "我是学习golang的人,hello word"
	str2 := "我是学习的人";
	xiangtong := strings.EqualFold(str , str2)
	fmt.Println("比较utf8:",xiangtong)
	//判断前缀
	hasPrefix := strings.HasPrefix(str , "我")
	fmt.Println("判断前缀:",hasPrefix)
	hasSuffix := strings.HasSuffix(str , "word")
	fmt.Println("判断后缀:",hasSuffix)
	//判断包含
	isContains := strings.Contains(str , "golang")
	fmt.Println("判断包含：",isContains)
	isContainsRune := strings.ContainsRune(str , 'g')
	fmt.Println("判断包含Rune：",isContainsRune)
	isContainsChars := strings.ContainsAny(str , "ord2")//只要有一个字符在，那就为true
	fmt.Println("判断包含chars：",isContainsChars)
	//
	count := strings.Count(str , "go")//返回字符串s中有几个不重复的sep子串。
	fmt.Println("数量count" , count)
	fmt.Println("indexByte:" , strings.IndexByte(str , 'g'))
	fmt.Println("IndexRune:" , strings.IndexRune(str , '是'))

	fmt.Println("Title:" , strings.Title(str))
	fmt.Println("ToTitle:" , strings.ToTitle(str))

	fmt.Println("Repeat:" ,strings.Repeat(str , 2) )
	//返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	fmt.Println("Replace:" , strings.Replace(str , "o" , "haha"  , 2))
	mapStr := strings.Map(func(r rune) rune {
		if r=='o'{
			return '改'
		}
		return r
	} , str)
	fmt.Println("Map:" , mapStr)

	str3 := "hello word golang good"
	fmt.Printf("Fields:%#v\n" , strings.Fields(str3))
	fmt.Printf("Split:%#v\n" , strings.Split(str3 , " "))
	fmt.Printf("SplitN:%#v\n" , strings.SplitN(str3 , " " , 2))

	//joinStr := make([]string , 0)
	//joinStr = append(joinStr , str , str2 , str3)
	joinStr := []string{str , str2 , str3}
	fmt.Printf("joinStr:%#v\n" , joinStr)
	fmt.Println("Join:" , strings.Join(joinStr , "----"))

	reader := strings.NewReader(str)
	fmt.Println("Read.Len" , reader.Len())
	num, _ := reader.Read([]byte(str))
	fmt.Println(num)


}
