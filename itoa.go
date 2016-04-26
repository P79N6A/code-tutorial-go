package main

import (
        "strconv"
        "fmt"
        "github.com/fvbock/trie"
        "net"
        "strings"
)
func ipToint(ip net.IP) uint64 {
        nip := ip.To4()
        fmt.Println(nip)
        fmt.Println(nip[0])
        fmt.Println(nip[1])
        fmt.Println(nip[2])
        fmt.Println(nip[3])
        return uint64(nip[0]) << 24 | uint64(nip[1]) << 16 |
        uint64(nip[2]) << 8 | uint64(nip[3])
}
func AddCIDR(cidr string) {
        ip,net,err := net.ParseCIDR(cidr)
        if err != nil {
                fmt.Errorf("Parse CIDR fail %s",cidr)
        } else {
                ipNum := ipToint(ip)
                mask,_ := net.Mask.Size()
                binIPStr := strconv.FormatUint(ipNum,2)
                fillField := strings.Repeat("0", 32 - len(binIPStr))
                binIPAll := fillField + binIPStr
                maskIP := binIPAll[:mask]
                fmt.Println(maskIP)
        }
}
func main() {
        fmt.Println(strconv.FormatInt(32,2))
        tree := trie.NewTrie()
        fmt.Println(tree.Dump())
        fmt.Println(tree.HasPrefix("123"))
        fmt.Println(tree.HasPrefix("23"))
        AddCIDR("1.2.3.4/24")

}
