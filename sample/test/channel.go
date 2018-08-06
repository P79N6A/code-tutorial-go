package main
import (
    "fmt"
)
func main() {
    // Here we make a channel of strings buffering up to 2 values.
    messages := make(chan string, 2)  // buffered channel with 2 values
    // Because this channel is buffered,
    // we can send these values into the channel without a corresponding concurrent receive.
    messages <- "buffered"
    messages <- "channel"
    messages <- "33"
    messages <- "xx"
    // Later we can receive these two values as usual.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}