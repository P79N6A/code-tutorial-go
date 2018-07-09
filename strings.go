package main

import (
        "fmt"
        "strings"
)

func main() {
        a := "abc123dswef"
        idx := strings.Index(a,"f")
        fmt.Println(len(a))
        fmt.Println(idx)
        fmt.Print(a[idx+1:])
}
