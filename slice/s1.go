package main

import "fmt"

func main(){

        s := make([]int,10)
        s[0]=1
        s[2]=3
        fmt.Println(s)
        j := append(s[:0], s[1:]...)
        fmt.Println(s)
        fmt.Println(j)
}
