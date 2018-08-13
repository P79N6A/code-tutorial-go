package main

import (
        "fmt"
)

type A struct {
        N string
}
func main(){
        var xxx []int
        for _,vv := range xxx {
                fmt.Println(vv)
        }

        s := make([]int,10)
        s[0]=1
        s[2]=3
        fmt.Println("XXXX")
        for _,v := range s {
                fmt.Println(v)
        }
        fmt.Println(len(s))
        fmt.Println("XXXX")
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
