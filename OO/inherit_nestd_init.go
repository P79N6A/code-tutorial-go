package main

import "fmt"

type Person struct {
        Name string
        Age int
}
type Man struct {
        Person
        Do string
}
func main() {
        m := &Man{
                Person:Person{
                        Name:"xxx",
                        Age:18,
                },
                Do:"hello",
        }
        fmt.Println(m)
}
