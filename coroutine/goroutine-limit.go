package main

import (
    "sync"
    "runtime"
    "time"
)

func main() {
    pool := NewWaitGroupWithPool(5)
    println(runtime.NumGoroutine())
    for i := 0; i < 30; i++ {
        pool.Add(1)
        go func() {
            time.Sleep(time.Second)
            println(runtime.NumGoroutine())
            pool.Done()
        }()
    }
    pool.Wait()
    println("XXXXXXXX")
    println(runtime.NumGoroutine())
    
}


type WaitGroupWithPool struct {
    queue chan int
    wg    *sync.WaitGroup
}

func NewWaitGroupWithPool(size int) *WaitGroupWithPool {
    if size <= 0 {
        size = 1
    }
    return &WaitGroupWithPool{
        queue: make(chan int, size),
        wg:    &sync.WaitGroup{},
    }
}

func (p *WaitGroupWithPool) Add(delta int) {
    for i := 0; i < delta; i++ {
        p.queue <- 1
    }
    for i := 0; i > delta; i-- {
        <-p.queue
    }
    p.wg.Add(delta)
}

func (p *WaitGroupWithPool) Done() {
    <-p.queue
    p.wg.Done()
}

func (p *WaitGroupWithPool) Wait() {
    p.wg.Wait()
    close(p.queue)
}
