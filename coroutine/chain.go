package main
import (
    "fmt"
    "time"
)
func ping(in <-chan string, out chan<- string) {
    for {
        msg := <- in
        fmt.Println("Ping" + msg)
        fmt.Println(len(in))
        fmt.Println(len(out))
        out <- msg
    }
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(in <-chan string, out chan<- string) {
    for {
        msg := <-in
        time.Sleep(5*time.Second)
        fmt.Println("Pong" + msg)
        out <- msg
    }
}
func outf(out chan string) {
    for {
        j := <-out
        time.Sleep(5*time.Second)
        fmt.Println("received job", j)
    }
}
func inf(in chan string) {
    for j := 0; j <= 100; j++ {
        in <- fmt.Sprintf("Get %d", j)
    }
}
func main() {
    in := make(chan string, 5)
    out := make(chan string, 5)
    out2 := make(chan string, 5)
    go inf(in)
    go ping(in, out)
    go pong(out, out2)
    go outf(out2)
   var input string
    fmt.Scanln(&input)
}
