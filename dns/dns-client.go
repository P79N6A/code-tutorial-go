package main

import (
    "github.com/miekg/dns"
    "net"
    "log"
    "fmt"
)

func main() {
    config := new(dns.ClientConfig)
    config.Servers = make([]string, 0)
    config.Search = make([]string, 0)
    config.Port = "53"
    config.Ndots = 1
    config.Timeout = 5
    config.Attempts = 2
    config.Servers = append(config.Servers,"8.8.8.8")

    c := new(dns.Client)

   // query := "www.a.shifen.com."
    query := "www.baidu.com"

    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(query), dns.TypeA)
    m.RecursionDesired = true

    r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
    if r == nil {
        log.Fatalf("*** error: %s\n", err.Error())
    }

    if r.Rcode != dns.RcodeSuccess {
        log.Fatalf(" *** invalid answer name %s after MX query for %s\n", query, query)
    }
    // Stuff must be in the answer section
    for _, a := range r.Answer {
        fmt.Printf("%v\n", a)
    }
}