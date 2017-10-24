package main

import (
        "net"
        "fmt"
	"strings"
)

func main() {
	var ip net.IP
	fmt.Println(ip)
	domains := strings.Split("a.", ".")
	fmt.Println(domains)
	fmt.Println(domains[1] == "")
	fmt.Println(len(domains))
	str := "1234"
	fmt.Println(str[:len(str)-1])
        i,in,_ := net.ParseCIDR("1.2.3.4/0")
        fmt.Println(i)
        fmt.Println(in)
        var ssss []string = nil
        fmt.Println(len(ssss))
}