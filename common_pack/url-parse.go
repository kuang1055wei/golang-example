package main

import (
	"fmt"
	"net/url"
)

func main() {
	//urlStr := "https://www.google.com/search?q=golang&sxsrf=18&source=hp&ei=kUZxYIAc&uact=5"
	urlStr :=  "postgres://user:pass@host.com:5432/path?k=v#f"
	u,_ := url.Parse(urlStr)
	fmt.Printf("%#v\n" , u)
	fmt.Println(u.User.Username())

	m,_ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
}
