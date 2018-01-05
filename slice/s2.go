package main

import "fmt"

func main() {
        ss := make(map[[]byte]int)
        ss[[]byte("xxx")] = 1
        if _,ok := ss[[]byte("xxx")];ok {
                fmt.Println("XXXXXXX")
        }
}
