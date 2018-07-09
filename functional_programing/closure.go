package main

import (
        "fmt"
        "sync"
)

func fn() func() int{
        var a int
        var mux sync.Mutex
        return func() int {
                mux.Lock()
                defer mux.Unlock()
                for i := 0;i < 1000;i++ {
                        a += 1
                }
                for i := 0;i < 1000;i++ {
                        a -= 1
                }
                a += 1
                return a
        }
}

func main() {
        g := fn()
        var wait sync.WaitGroup
        for i := 0;i < 10;i++ {
                wait.Add(1)
                go func() {
                        fmt.Println(g())
                        wait.Done()
                }()
        }
        wait.Wait()
}

