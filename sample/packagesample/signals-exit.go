/*
	Signals
	Sometimes we’d like our Go programs to intelligently handle Unix signals.
	For example, we might want a server to gracefully shutdown when it receives a SIGTERM,
	or a command-line tool to stop processing input if it receives a SIGINT.
	Here’s how to handle signals in Go with channels
	Exit
	Use os.Exit to immediately exit with a given status
*/
package psample

import "fmt"
import "os"
import "os/signal"
import "syscall"
/*
When we run this program it will block waiting for a signal. By typing ctrl-C (which the terminal shows as ^C) we can send a SIGINT signal, causing the program to print interrupt and then exit.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
*/
func Signal() {
//Go signal notification works by sending os.Signal values on a channel. We’ll create a channel to receive these notifications (we’ll also make one to notify us when the program can exit).
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)
//signal.Notify registers the given channel to receive notifications of the specified signals.
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
//This goroutine executes a blocking receive for signals. When it gets one it’ll print it out and then notify the program that it can finish.
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()
//The program will wait here until it gets the expected signal (as indicated by the goroutine above sending a value on done) and then exit.
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
/*
Note that unlike e.g. C, Go does not use an integer return value from main to indicate exit status. If you’d like to exit with a non-zero status you should use os.Exit.
If you run exit.go using go run, the exit will be picked up by go and printed.
$ go run exit.go
exit status 3
By building and executing a binary you can see the status in the terminal.
$ go build exit.go
$ ./exit
$ echo $?
3
*/
func Exits() {
//defers will not be run when using os.Exit, so this fmt.Println will never be called.
    defer fmt.Println("!")
//Exit with status 3.
    os.Exit(3)
}