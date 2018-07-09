package main
import (
	"github.com/miekg/dns"
	"fmt"
	"net"
	"reflect"
)
func newA(rr string) *dns.A           {
        r, _ := dns.NewRR(rr);
        if r == nil {
                fmt.Println("eeeeee")
        }

        if r.(*dns.A) == nil {
                fmt.Println("XXXXX")
        }
        return r.(*dns.A)
}
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
func main() {
	aa := net.ParseIP("xx.1.2.3")
	fmt.Println(aa==nil)
	cname := newCNAME("backend.in.skydns.test. IN CNAME ipaddr.skydns.test.")
	fmt.Println(cname.Target)
	a := newA("a.miek.nl. 3600 IN A 176.58.119.54")
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(&dns.A{}))
	fmt.Println(a.A.String())

        soa := newSOA("jerry.com.      IN      SOA    ns1.tom.com. webmaster.ewcache.com. 1077295087 10800 3600 604800 600")
        fmt.Println(soa)

	ip := net.ParseIP("1.2.3.4")
	ip4 := ip.To4()
	fmt.Println(ip4[1])
	fmt.Println(len(ip4))
        lable,ok:=dns.IsDomainName("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXssssssssssssssssssssssssssXXXX.com")
        fmt.Println(lable)
        fmt.Println(ok)
        //aaa := newA("iiiiiiiiiiiiiiiiiiioooiioooooiii33333333333333333333333rrr9rrrrr.COM. 600 IN A 11.1.1.1")
        //aaa := newA("wekrjkekrekrekrrrrrrrrrrr3333333333333333333333333333333333333333333333333333333333333333333333333333333333rrr9rrrrr.COM. 600 IN A 11.1.1.1")
        //fmt.Println(aaa)
}
