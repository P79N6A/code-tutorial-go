package main

import (
    "fmt"
)

func main() {
    messages := make(chan string) //如果不加缓存的话，就全部会选择defalut
    //messages := make(chan string, 1) //加了缓存的话，会选择对应的
    signals := make(chan bool)

    messages <- "test" //因为该channels没有缓存，而其又赋值了，会到时死锁。

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // go func() {
    msg := "hi world"
    // }()
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}

