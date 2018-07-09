package main
import (
	"github.com/miekg/dns"
	"fmt"
	"net"
	"strings"
	"reflect"
)
func newCNAME(rr string) *dns.CNAME   { r, _ := dns.NewRR(rr); return r.(*dns.CNAME) }
func newA(rr string) *dns.A           { r, _ := dns.NewRR(rr); return r.(*dns.A) }
func main() {
	aa := net.ParseIP("xx.1.2.3")
	fmt.Println(aa==nil)
	cname := newCNAME("backend.in.skydns.test. IN CNAME ipaddr.skydns.test.")
	fmt.Println(cname.Target)
	a := newA("a.miek.nl. 3600 IN A 176.58.119.54")
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(&dns.A{}))
	fmt.Println(a.A.String())

	slist := []string{"x"}
	fmt.Println(strings.Join(slist,","))

	ip := net.ParseIP("1.2.3.4")
	ip4 := ip.To4()
	fmt.Println(ip4[1])
	fmt.Println(len(ip4))
        lable,ok:=dns.IsDomainName("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXssssssssssssssssssssssssssXXXX.com")
        fmt.Println(lable)
        fmt.Println(ok)
}
