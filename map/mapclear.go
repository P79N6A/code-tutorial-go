package main

import "fmt"

var a map[string]string
var b map[string]string

func main() {
    fmt.Println(a==nil)
    a = make(map[string]string)
    fmt.Println(a==nil)
    b=a
    a["hello"]="world"
    a = nil
    fmt.Println(a)
}
