package main

import "fmt"

type A struct {
        N string
}
func main(){

        s := make([]int,10)
        s[0]=1
        s[2]=3
        fmt.Println(s)
        j := append(s[:0], s[1:]...)
        fmt.Println(s)
        fmt.Println(j)

        fmt.Println("xxxx")
        var sl []*A = nil
        for _,a := range sl {
                fmt.Println(a)
        }
        fmt.Println("xxxx")
}
