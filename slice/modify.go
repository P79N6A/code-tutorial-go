package main

import "fmt"

func main() {
        s := make([]string,0)
        s = append(s,"aaa")
        s = append(s,"bbb")
        s = append(s,"ccc")
        fmt.Println(s)
        for _,a := range s {
                a = a + "xx"
        }

        fmt.Println(s)
}
