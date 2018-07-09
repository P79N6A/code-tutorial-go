package main

import "fmt"
import "github.com/miekg/dns"

func newCNAME(rr string) *dns.CNAME   { r, _ := dns.NewRR(rr); return r.(*dns.CNAME) }
func main() {
        cname := newCNAME("backend.com IN CNAME ipaddr.skydns.test.")
        fmt.Println(cname.String())
}
