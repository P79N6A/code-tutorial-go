package main

import (
        "fmt"
        "net"
        "strings"
        "ipset/src/common"
        "math"
)

func GetIPsFromCIDR(cidr string) (error, net.IP, net.IP) {
        _,ipnet,err := net.ParseCIDR(cidr)
        if err != nil {
                return err,nil,nil
        }
        ip4 := ipnet.IP.To4()
        if ip4 == nil {
                return fmt.Errorf("NoV4"),nil,nil
        }
        newip := make([]byte,4)
        for i := 0;i < 4 ;i ++ {
                newip[i] = []byte(ip4)[i] ^ (^[]byte(ipnet.Mask)[i])
        }
        fmt.Printf("%s,%s\n",ip4.String(),net.IP(newip).String())
        return nil,ip4,net.IP(newip)
}
func main() {
        GetIPsFromCIDR("14.208.0.0/12")
        GetIPsFromCIDR("39.170.0.0/15")
        GetIPsFromCIDR("39.172.0.0/15")
        GetIPsFromCIDR("39.174.0.0/18")
        GetIPsFromCIDR("39.175.0.0/32")
        GetIPsFromCIDR("39.175.191.0/24")
        GetIPsFromCIDR("39.175.192.0/18")
        fmt.Println("XXXXXXXXXXXXXXXXX")
        str := "39.170.0.0/15;39.172.0.0/14"
        for _,s := range strings.Split(str,";") {
                GetIPsFromCIDR(s)
        }
        cidrs,_ := common.GenCIDRFromIP(net.ParseIP("119.38.219.228"),net.ParseIP("119.38.219.230"))
        fmt.Println(cidrs)
        fmt.Println(math.MaxInt32)
}
