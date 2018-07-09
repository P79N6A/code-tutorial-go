package main

import "fmt"

func main() {
        done := make(chan bool,1)
        for i:=0; i<3 ;i++ {
                amap := f(i)
                go afunc(amap, done)
        }
        for i:=0; i<3; i++ {
                <- done
        }
}
func f(i int) map[string]int {
        amap := make(map[string]int)
        amap["a"] = 0
        amap["b"] = 1
        return nil
}
func afunc(amap map[string]int, done chan bool) {
        fmt.Println(amap)
        done <- true
}