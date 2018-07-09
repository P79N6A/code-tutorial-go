package main

import (
        "time"
        "fmt"
)

func main() {
        var a time.Duration = 100
        a = a * 98 / 100
        fmt.Println(a)
}
