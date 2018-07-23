package main

import "fmt"

func main() {
    x := 9
    fmt.Println(^x+1)
    fmt.Println(^(x-1))
    fmt.Println(x&(-x))

}
