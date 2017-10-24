package main

import (
        "net"
        "fmt"
)

func main() {
        ip,ipn,_ := net.ParseCIDR("195.51.142.0/15")
        fmt.Println(ip)
        fmt.Println(ipn)
}
