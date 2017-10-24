package main

import "fmt"

type Container struct {
        E interface{}
}
func main() {
        c := new(Container)
        c.E = "xxxxx"
        if str,ok := c.E.(string);ok {
                fmt.Println("str")
                fmt.Println(str)
        }
}
