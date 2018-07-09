package main

import "fmt"

type Element struct {
    val int
    name string
}
type Elements []Element

func main() {
    e1 := Element{1,"xx"}
    e2 := Element{1,"xx"}
    e3 := &Element{2,"xx"}
    if e1 == e2 {
        fmt.Println("equall...")
    }
    fmt.Println(e3)

}
