package main

import "strings"
import (
        "github.com/miekg/dns"
        "fmt"
)

func DomainLevel(domain string) (int,string) {
        domain = strings.ToLower(domain)
        domain = dns.Fqdn(domain)
        level := 0
        if domain[0] == '.' {
                level = -1
        }
        for _,d := range domain {
                if d == '.' {
                        level += 1
                }
        }
        return level,domain
}
func main() {
        i,j := DomainLevel("")
        fmt.Println(i,j)
        i,j = DomainLevel("com.")
        fmt.Println(i,j)
        i,j = DomainLevel(".com.")
        fmt.Println(i,j)
        i,j = DomainLevel("ss.com.")
        fmt.Println(i,j)
        fmt.Println(dns.IsDomainName(""))
}
