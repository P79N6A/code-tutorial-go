package main

import "fmt"

func main() {
    a := []string{"1","2"}
    b := []string{"3","2"}
    for _,x := range append(a,b...) {
        fmt.Println(x)
    }

}
