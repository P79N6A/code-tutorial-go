package main

import (
    "github.com/miekg/dns"
    "net"
    "log"
    "fmt"
    "time"
)
/*
Cname:src.iduwpui.qiniudns.com.	572	IN	CNAME	nb-gate-io-src.qiniu.com.
A:nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.5
A:nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.23
A:nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.6
A:nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.7
A:nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.8
A:nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.21
;; opcode: QUERY, status: NOERROR, id: 30362
;; flags: qr ra; QUERY: 1, ANSWER: 7, AUTHORITY: 2, ADDITIONAL: 12

;; QUESTION SECTION:
;src.iduwpui.qiniudns.com.	IN	 A

;; ANSWER SECTION:
src.iduwpui.qiniudns.com.	572	IN	CNAME	nb-gate-io-src.qiniu.com.
nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.5
nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.23
nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.6
nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.7
nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.8
nb-gate-io-src.qiniu.com.	171	IN	A	101.71.123.21

;; AUTHORITY SECTION:
qiniu.com.	30674	IN	NS	ns3.dnsv5.com.
qiniu.com.	30674	IN	NS	ns4.dnsv5.com.

;; ADDITIONAL SECTION:
ns3.dnsv5.com.	95697	IN	A	121.51.2.171
ns3.dnsv5.com.	95697	IN	A	125.39.213.190
ns3.dnsv5.com.	95697	IN	A	180.153.10.169
ns3.dnsv5.com.	95697	IN	A	182.140.167.191
ns3.dnsv5.com.	95697	IN	A	183.60.57.192
ns3.dnsv5.com.	95697	IN	A	183.60.59.217
ns3.dnsv5.com.	95697	IN	A	184.105.206.63
ns3.dnsv5.com.	95697	IN	A	112.90.143.36
ns3.dnsv5.com.	95697	IN	A	115.236.151.180
ns3.dnsv5.com.	95697	IN	A	117.135.170.109
ns3.dnsv5.com.	95697	IN	A	119.28.48.224
ns3.dnsv5.com.	95697	IN	A	119.167.195.9

*/
func main() {
    config := new(dns.ClientConfig)
    config.Servers = make([]string, 0)
    config.Search = make([]string, 0)
    config.Port = "53"
    config.Ndots = 1
    config.Timeout = 5
    config.Attempts = 2
//    localdns := "127.0.0.1"
	localdns := "202.96.64.68"
    //localdns := "114.114.114.114"
    config.Servers = append(config.Servers, localdns)

    c := &dns.Client{
        DialTimeout:time.Second,
    }

   // query := "www.a.shifen.com."
    query := "src.iduwpui.qiniudns.com"

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