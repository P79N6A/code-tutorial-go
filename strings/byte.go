package main

import "fmt"

func main() {
    b := make([]byte,10)
    for i:=0;i<len(b);i++ {
        b[i]=' '
    }
    fmt.Println(string(b))

}
