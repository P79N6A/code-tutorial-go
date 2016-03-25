package main

import (
    "github.com/miekg/dns"
    "net"
    "log"
    "fmt"
    "time"
)

func main() {
    config := new(dns.ClientConfig)
    config.Servers = make([]string, 0)
    config.Search = make([]string, 0)
    config.Port = "53"
    config.Ndots = 1
    config.Timeout = 5
    config.Attempts = 2
    localdns := "127.0.0.1"
    //localdns := "114.114.114.114"
    config.Servers = append(config.Servers, localdns)

    c := &dns.Client{
        DialTimeout:time.Millisecond,
    }

   // query := "www.a.shifen.com."
    query := "www.baidu.com"

    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(query), dns.TypeA)
    m.RecursionDesired = false

    r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
    if r == nil {
        log.Fatalf("*** error: %s\n", err.Error())
    }

    if r.Rcode != dns.RcodeSuccess {
        log.Fatalf(" *** invalid answer name %s after MX query for %s\n", query, query)
    }
    // Stuff must be in the answer section
    for _, a := range r.Answer {
        if a.Header().Rrtype == dns.TypeA {
            fmt.Printf("A:%v\n", a)
        } else {
            fmt.Printf("Cname:%v\n",a)
        }
    }

    fmt.Println(r.String())
}