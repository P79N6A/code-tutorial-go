package main

import (
    "fmt"
)

func main() {
    // messages := make(chan string)    //如果不加缓存的话，就全部会选择defalut
    messages := make(chan string, 1) //加了缓存的话，会选择对应的
    signals := make(chan bool)

    // messages <- "test"

    select {
    case msg := <-messages:
        fmt.Println("received message", msg) //因为messages目前本身还没有值，因此选择default执行
    default:
        fmt.Println("no message received")
    }

    // go func() {
    msg := "hi world"
    // }()
    select {
    case messages <- msg:
        fmt.Println("sent message", msg) //因为channels有缓存，所以这里的msg发送到 channels messages 能处理，不会被阻塞住
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg) //因为messages已经有值了，所以会选择这个case执行
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}

