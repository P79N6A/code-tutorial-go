package main

import "fmt"

type Interface interface {
    hello()
}
type X struct {

}
func (x *X)hello() {
    fmt.Println("hello")
}

func foo(xx Interface) {
    xx.hello()
}
func main() {
    foo(X{})
}
