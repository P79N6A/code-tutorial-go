/*
	Timers:
		We often want to execute Go code at some point in the future,
		or repeatedly at some interval.
		Go’s built-in timer and ticker features make both of these tasks easy.
		We’ll look first at timers and then at tickers.
	Tickers:
		Timers are for when you want to do something once
		in the future - tickers are for when you want to do something repeatedly at regular intervals.
		Here’s an example of a ticker that ticks periodically until we stop it.
		
*/
package main

import (
	"fmt"
	"time"
)
/*
The first timer will expire ~2s after we start the program,
but the second should be stopped before it has a chance to expire.
$ go run timers.go
Timer 1 expired
Timer 2 stopped
*/
func Timers() {
	// Timers represent a single event in the future.
	// You tell the timer how long you want to wait,
	// and it provides a channel that will be notified at that time.
	// This timer will wait 2 seconds.
    timer1 := time.NewTimer(time.Second * 2)
	// The <-timer1.C blocks on the timer’s channel C
	// until it sends a value indicating that the timer expired.
    <-timer1.C
    fmt.Println("Timer 1 expired")
	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is
	// that you can cancel the timer before it expires. Here’s an example of that.
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
func Tickers() {
	// Tickers use a similar mechanism to timers:
	// a channel that is sent values.
	// Here we’ll use the range builtin on the channel
	// to iterate over the values as they arrive every 500ms.
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()
	// Tickers can be stopped like timers.
	// Once a ticker is stopped it won’t receive any more values on its channel.
	// We’ll stop ours after 1500ms.
    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}

func main() {
    Timers()
    Tickers()
}