package main

import (
        "fmt"
        "strings"
        "net"
        "time"
        "net/url"
)

func main() {
        aa := strings.Split("com.",".")
        fmt.Println(aa)
        fmt.Println(len(aa))
        var ip net.IP
        fmt.Println(ip)
        ttt := int32(time.Millisecond)
        fmt.Println(ttt)
        var ssss  []string = nil
        for _,v := range ssss {
                fmt.Println(v)
        }
        fmt.Println("-------")
        u,err := url.Parse("xxx")
        fmt.Println(err)
        fmt.Println(u)
        //////////////
        ips := "200.1.1.1:-100"

        addr := strings.Split(ips,":")
        addr0 := strings.TrimSpace(strings.Join(addr[:len(addr)-1],":"))
        fmt.Println(addr0)
        fmt.Println(strings.TrimSpace(addr[len(addr)-1]))
}
