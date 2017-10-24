package main

import "fmt"

func main(){
        fmt.Println(ff())
}
func ff() string {
        str := "xxx"
        defer func(pstr *string) {
                 if err:=recover();err!=nil{
                         fmt.Println("a")
                         *str = "jfwiefjoei"
                }

        }(&str)
        panic(33)
        return str
}

func f(){
        fmt.Println("a")
        panic(55)
        fmt.Println("b")
        fmt.Println("f")
}
