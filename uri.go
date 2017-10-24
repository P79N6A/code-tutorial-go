package main

import (
	"net/url"
	"fmt"
)
func main() {
	uri := "mysql://user:password@127.0.0.1:3306/dbname"
	u,e := url.Parse(uri)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.Fragment)
	fmt.Println(u.Host)
	fmt.Println(u.Path[1:])
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
        var aaaa []string
        fmt.Println(len(aaaa))
}
