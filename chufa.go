package main

import (
        "fmt"
        "net"
)

func main() {
        var a,b,c uint64 = 15,999,63
        var aa,bb,cc uint64 = 1500,999,6300
        fmt.Println(a*b/c)
        fmt.Println(aa*bb/cc)
        ipip := net.ParseIP("2001::9d38:953c:427:b791:c522:6b5d")
        fmt.Println(ipip.To16())
}
