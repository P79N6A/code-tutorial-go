package main

import (
        "net"
        "fmt"
        "ipset/src/common"
)

func main() {
        ip,ipnet,_ := net.ParseCIDR("1.2.3.4/25")
        fmt.Println(ip)
        fmt.Println(ipnet.IP)
        fmt.Println([]byte(ipnet.Mask))
        //fmt.Println(int(ipnet.Mask) & 0xf0000000)
        ipu := common.IP2Uint(net.IP(ipnet.Mask)) & 0xffffff00
        ipnet.Mask = net.IPMask(common.Uint2IP(ipu))
        fmt.Println(ipnet.String())
}
