package main

import "fmt"

/*
https://tenfy.cn/2017/09/20/golang-unbuffered-channel/

If the capacity is zero or absent,
the channel is unbuffered and communication succeeds only when both a sender and receiver are ready.
Otherwise, the channel is buffered and communication succeeds without blocking
if the buffer is not full (sends) or not empty (receives).

其实官网说得很明白。如果make时，容量传0或者不传的话，即创建一个没有buffer的 channel，
它们只有发送方和接收方都准备好了，才能通信成功。不然，只要buffer不满就 可以立即发送，或者不空就可以立即接收。

也就是说，如果是没有buff的，必须sender和receiver在两个goroutine，并且发送要等接收端准备好。否则block住。 下边的例子用select当然不会block，那就只能执行default。
因为下边的例子没有任何其他goroutine存在

有buff是不会blcok发送端的。buff是1，那发送端最多发送一个

*/
func main() {
    messages := make(chan string)
    signals := make(chan bool)

    /**
    这里是一个非阻塞 receive。如果在 messages 上的值是可用的，那 select 将 <-messages 的值带上，执行 <-messages 下面的println语句。如果不是，它将立即带上 default 的值，执行 default 下面的println语句
    **/
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }
    /**
    一个非阻塞 send 的类似工作
    **/
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    /**
    我们可以用在 default 之上使用多个 cases 来实现一个非阻塞的多路 select。在这里我们尝试在 messages 和 signals 上实现非阻塞 receives。
    **/

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
