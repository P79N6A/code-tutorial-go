package main

import (
        "fmt"
        "reflect"
)

type HaHa interface {
        Hello()
}
type OneHaHa struct {
}
func (*OneHaHa)Hello()  {
        fmt.Println("onehaha")
}
type TwoHaHa struct {
}
func (*TwoHaHa)Hello()  {
        fmt.Println("TwoHaHa")
}
func main() {
        var haha HaHa
        haha = new(TwoHaHa)
        fmt.Println(reflect.TypeOf(haha).String())
        fmt.Println(reflect.TypeOf(haha).Elem())
}
