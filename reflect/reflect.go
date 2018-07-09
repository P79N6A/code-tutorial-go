package main

import "fmt"
import "reflect"

type T struct {}

func (t *T) Foo() *T {
        return &T{}
}
type Finterface interface {
        Foo() string
}

type Person struct {
        name string
        aget int
}
func (p Person)String(int) int {
        return p.aget
}
func f1 () {
        kelu := Person{"kelu", 25}
        t := reflect.TypeOf(kelu)
        n := t.NumMethod()
        for i := 0; i < n; i++ {
                fmt.Println(t.Method(i).Name)
                fmt.Println(t.Method(i).Type)
        }
}
func f2() {
        var t interface{}
        t = &T{}

        n := reflect.ValueOf(t).MethodByName("Foo").IsValid()
        fmt.Println(n)
        ret := reflect.ValueOf(t).MethodByName("Foo").Call(nil)
        fmt.Println(ret)
        for _,v := range ret {
                fmt.Println(v.String())
        }
}
func main() {
        f2()
}