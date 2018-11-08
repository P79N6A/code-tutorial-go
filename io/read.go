package main

import "fmt"

func read(b []byte) {
    x := []byte("abcdefg")
    copy(b,x)
}
func main() {

    b := make([]byte,10)
    read(b)
    fmt.Println(string(b))
}
