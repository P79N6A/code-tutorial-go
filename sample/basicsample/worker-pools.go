/*
	Worker Pools:
		In this example we’ll look at how to implement a worker pool using goroutines and channels.
		
*/
package main

import (
	"fmt"
	"time"
)
/*
Our running program shows the 9 jobs being executed by various workers.
 The program only takes about 3 seconds despite doing about 9 seconds of total work
  because there are 3 workers operating concurrently.
$ time go run worker-pools.go 
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
worker 1 processing job 4
worker 2 processing job 5
worker 3 processing job 6
worker 1 processing job 7
worker 2 processing job 8
worker 3 processing job 9
real	0m3.149s
*/
// Here’s the worker, of which we’ll run several concurrent instances.
// These workers will receive work on the jobs channel and
// send the corresponding results on results.
// We’ll sleep a second per job to simulate an expensive task.
func oneworker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}
func WorkerPools() {
	// In order to use our pool of workers we need to send them work and collect their results.
	// We make 2 channels for this.
    jobs := make(chan int, 100)
    results := make(chan int, 100)
	// This starts up 3 workers, initially blocked because there are no jobs yet.
    for w := 1; w <= 3; w++ {
        go oneworker(w, jobs, results)
    }
	// Here we send 9 jobs and then close that channel to indicate that’s all the work we have.
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
	// Finally we collect all the results of the work.
    for a := 1; a <= 9; a++ {
        fmt.Println(<-results)
    }
}
func main()  {
    WorkerPools()
}