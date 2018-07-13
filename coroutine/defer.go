package main

import "fmt"

func main() {
    defer func() {fmt.Println("xxxxx")}()
    defer func() {fmt.Println("yyyyy")}()
}
