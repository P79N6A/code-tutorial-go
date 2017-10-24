package main

import (
        "time"
        "fmt"
)

func work(s *[]string) {
        for {
           time.Sleep(time.Millisecond*10)
                *s = append(*s,"a")
        }
}
func main() {
        s := make([]string,0)
        for i := 0; i < 5; i ++ {
                go work(&s)
        }
        for {
                fmt.Println(s)
                time.Sleep(time.Second)
        }
}
