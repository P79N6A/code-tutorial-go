package main

import (
        "net"
        "fmt"
        "ipset/src/common"
)

func main() {
        ip,ne,_ := net.ParseCIDR("1.2.3.4/19")
        fmt.Println(ne.Mask.Size())
        fmt.Println(ip)
        fmt.Println([]byte(ne.IP))
        fmt.Println([]byte(ne.Mask))
        fmt.Println(ne.Contains(net.ParseIP("1.2.3.44")))
        fmt.Println([]byte(net.ParseIP("1.2.3.4").To4()))
        fmt.Println(string([]byte{0,0,0,0,0,0,0,1}))
        var addr net.IP
        fmt.Println(string(addr))

        ip1 := net.ParseIP("183.207.112.0")
        ip3 := net.ParseIP("183.207.113.255")
        cidrs,_ := common.GenCIDRFromIP(ip1,ip3)
        fmt.Println(cidrs)
        fmt.Println(common.IP2Uint(net.ParseIP("113.215.248.0")))
        fmt.Println(common.IP2Uint(net.ParseIP("113.215.255.255")))
        fmt.Println(net.ParseIP("fe80::ed57:2e8a:ea71:9ed0"))
}
