package main

import "fmt"

func main() {
    
    msgs := make(chan string, 2)
    
    msgs <- "buffered"
    msgs <- "channel"

    fmt.Println(<- msgs)
    fmt.Println(<- msgs)

}