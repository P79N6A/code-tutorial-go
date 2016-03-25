package main

import (
    "github.com/miekg/dns"
    "time"
    "fmt"
    "net"
)
// good sample....
// https://github.com/mindreframer/golang-devops-stuff/blob/bb6c6c16ff0ae7892cd0ed0715b8a6b769052f9f/src/github.com/skynetservices/skydns/server_test.go


func newA(rr string) *dns.A           { r, _ := dns.NewRR(rr); return r.(*dns.A) }
func newAAAA(rr string) *dns.AAAA     { r, _ := dns.NewRR(rr); return r.(*dns.AAAA) }
func newCNAME(rr string) *dns.CNAME   { r, _ := dns.NewRR(rr); return r.(*dns.CNAME) }
func newSRV(rr string) *dns.SRV       { r, _ := dns.NewRR(rr); return r.(*dns.SRV) }
func newSOA(rr string) *dns.SOA       { r, _ := dns.NewRR(rr); return r.(*dns.SOA) }
func newNS(rr string) *dns.NS         { r, _ := dns.NewRR(rr); return r.(*dns.NS) }
func newDNSKEY(rr string) *dns.DNSKEY { r, _ := dns.NewRR(rr); return r.(*dns.DNSKEY) }
func newRRSIG(rr string) *dns.RRSIG   { r, _ := dns.NewRR(rr); return r.(*dns.RRSIG) }
func newNSEC3(rr string) *dns.NSEC3   { r, _ := dns.NewRR(rr); return r.(*dns.NSEC3) }
func newPTR(rr string) *dns.PTR       { r, _ := dns.NewRR(rr); return r.(*dns.PTR) }
func newTXT(rr string) *dns.TXT       { r, _ := dns.NewRR(rr); return r.(*dns.TXT) }


var server = &dns.Server{Addr: ":53", Net: "udp"}
func main() {
   // server.TsigSecret = map[string]string{"axfr.": "so6ZGir4GPAqINNh9U5c3A=="}
    dns.HandleFunc(".", handleRequest)
    go func() {
        err := server.ListenAndServe()
        if err != nil {
            fmt.Printf("Error %s", err.Error())
            return
        }
    }()
    for true {
        time.Sleep(time.Second)
    }
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
    m := new(dns.Msg)
    fmt.Printf("get query %s", r.Question[0].Name)
    fmt.Printf("Query: %s",r.String())
    // remote request client ip...

    m.SetReply(r)

    m.Answer = []dns.RR{
        &dns.CNAME{
            Hdr:dns.RR_Header{Name: m.Question[0].Name, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 120},
            Target:"cname.domain.com.",// target must end with .!!!!!!
        },
        newCNAME("backend.in.skydns.test. IN CNAME ipaddr.skydns.test."),
        &dns.A{
            //Hdr:dns.RR_Header{Name: "cname.domain.com", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 120},
            Hdr:dns.RR_Header{Name: m.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 120},
            A:net.ParseIP("1.1.1.1"),
        },

        &dns.A{
            //Hdr:dns.RR_Header{Name: "cname.domain.com", Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 120},
            Hdr:dns.RR_Header{Name: m.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 120},
            A:net.ParseIP("1.1.1.1"),
        },
    }
    /*m.Answer = []dns.RR{
        newCNAME("backend.in.skydns.test. IN CNAME ipaddr.skydns.test."),
        newA("ipaddr.skydns.test. IN A 172.16.1.1"),
        newA("ipaddr.skydns.test. IN A 172.16.1.2"),
    }*/

    m.Ns = make([]dns.RR,2)
    m.Ns[0] = newSOA("qiniudns.com.           60      IN      SOA     ns3.dnsv5.com. enterprise3dnsadmin.dnspod.com. 1457923437 3600 180 1209600 180")

    m.Extra = make([]dns.RR, 1)
    m.Extra[0] = &dns.TXT{Hdr: dns.RR_Header{Name: m.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0}, Txt: []string{"Hello example"}}

    w.WriteMsg(m)
}

