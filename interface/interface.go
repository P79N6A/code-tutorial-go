package main

import "fmt"

func Itf(itf interface{}) {
        str,ok := itf.(string)
        if ok {
                fmt.Println(str)
        }
}

func main() {
        Itf(nil)
}
