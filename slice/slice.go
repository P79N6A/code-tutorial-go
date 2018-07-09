package main

import "fmt"

func f(p *[]byte) {
        return
}
func main() {
        path := []byte("12345")
        fmt.Println(path[:len(path)-1])
        mp := make(map[string]string)
        mp[""] = "xx"
        mp["1"] = "yy"
        fmt.Println(mp)
        fmt.Println(mp[""])
}
