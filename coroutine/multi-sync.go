package main

import (
        "fmt"
        "time"
)

func work() {
        fmt.Println("aaa")
}

func main() {
        dones := make([]chan error,0)
        for i := 0;i < 10000;i++ {
                done := make(chan error)
                dones = append(dones,done)
                go func(done chan error) {
                        work()
                        time.Sleep(time.Second * 10)
                        done <- nil
                }(done)
        }
        fmt.Println("syncsync")
        for _,d := range dones {
                err := <- d
                if err != nil {
                        fmt.Println(err)
                }
        }
        fmt.Println("endendend")
}
