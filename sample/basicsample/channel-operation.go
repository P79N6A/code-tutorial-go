/*
	Select:
		Go’s select lets you wait on multiple channel operations. 
		Combining goroutines and channels with select is a powerful feature of Go.
	Timeouts:
		Timeouts are important for programs that connect to external resources or
		that otherwise need to bound execution time.
		Implementing timeouts in Go is easy and elegant thanks to channels and select.
	Non-Blocking Channel Operations:
		Basic sends and receives on channels are blocking. However,
		we can use select with a default clause to implement non-blocking sends,
		receives, and even non-blocking multi-way selects.
	Closing Channels:
		Closing a channel indicates that no more values will be sent on it.
		This can be useful to communicate completion to the channel’s receivers.
	Range over Channels:
		In a previous example we saw how for and range provide iteration over basic data structures.
		We can also use this syntax to iterate over values received from a channel.
		
*/
package bsample

import (
	"fmt"
	"time"
)
/*
	$time go run select.go 
	received one
	received two
	real	0m2.245s

	Note that the total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.
*/
func SelectChannel() {
	// For our example we’ll select across two channels.
    c1 := make(chan string)
    c2 := make(chan string)
	// Each channel will receive a value after some amount of time,
	// to simulate e.g. blocking RPC operations executing in concurrent goroutines.
    go func() {
        time.Sleep(time.Second * 5)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 3)
        c2 <- "two"
    }()
	// We’ll use select to await both of these values simultaneously,
	// printing each one as it arrives.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}


/*
Running this program shows the first operation timing out and the second succeeding.
$ go run timeouts.go 
timeout 1
result 2
Using this select timeout pattern requires communicating results over channels.
This is a good idea in general because other important Go features are based on channels and select. 
*/
func TimeoutChannel() {
	// For our example, suppose we’re executing an external call
	// that returns its result on a channel c1 after 2s.
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()
	// Here’s the select implementing a timeout.
	// res := <-c1 awaits the result and <-Time.
	// After awaits a value to be sent after the timeout of 1s.
	// Since select proceeds with the first receive that’s ready,
	// we’ll take the timeout case if the operation takes more than the allowed 1s.
    select {
	    case res := <-c1:
	        fmt.Println(res)
	    case <-time.After(time.Second * 1):
	        fmt.Println("timeout 1")
    }
	// If we allow a longer timeout of 3s,
	// then the receive from c2 will succeed and we’ll print the result.
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}


func NonBlockChannel() {
    messages := make(chan string)
    signals := make(chan bool)
	// Here’s a non-blocking receive.
	// If a value is available on messages then select will take the <-messages case with that value.
	// If not it will immediately take the default case.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }
	// A non-blocking send works similarly.
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }
	// We can use multiple cases above the default clause to implement a multi-way non-blocking select.
	// Here we attempt non-blocking receives on both messages and signals.
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}



/*
$ go run closing-channels.go 
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs
*/
// In this example we’ll use a jobs channel to communicate work
// to be done from the main() goroutine to a worker goroutine.
// When we have no more jobs for the worker we’ll close the jobs channel.
func CloseChannel() {
    jobs := make(chan int, 5)
    done := make(chan bool)
	// Here’s the worker goroutine.
	// It repeatedly receives from jobs with j,more := <-jobs.
	// In this special 2-value form of receive,
	// the more value will be false if jobs has been closed and
	// all values in the channel have already been received.
	// We use this to notify on done when we’ve worked all our jobs.
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()
	// This sends 3 jobs to the worker over the jobs channel, then closes it.
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")
	// We await the worker using the synchronization approach we saw earlier.
    <-done
}



func RangeOverChannel() {
	// We’ll iterate over 2 values in the queue channel.
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
	// This range iterates over each element as it’s received from queue.
	// Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	// If we didn’t close it we’d block on a 3rd receive in the loop.
    for elem := range queue {
        fmt.Println(elem)
    }
}