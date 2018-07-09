package main

import (
        "net"
        "fmt"
)

func main() {
        xxx := make(map[string]bool)
        ipa := net.ParseIP("1.2.3.4")
        xxx[string(ipa)] = true
        fmt.Println(xxx)

        var yyy map[string]string = nil
        for k, v := range yyy {
                fmt.Println(k)
                fmt.Println(v)
        }
        fmt.Println("yyyy")
}
