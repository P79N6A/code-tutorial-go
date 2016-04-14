package main

import (
	"github.com/miekg/dns"
	"strings"
	"fmt"
)
func main() {
	zone := `$ORIGIN .
$TTL 3600       ; 1 hour
name                    IN SOA  a6.nstld.com. hostmaster.nic.name. (
                                203362132  ; serial
                                300        ; refresh (5 minutes)
                                300        ; retry (5 minutes)
                                1209600    ; expire (2 weeks)
                                300        ; minimum (5 minutes)
                                )
$TTL 10800      ; 3 hours
name.	10800	IN	NS	name.
               IN       NS      g6.nstld.com.
               7200     NS      h6.nstld.com.
             3600 IN    NS      j6.nstld.com.
             IN 3600    NS      k6.nstld.com.
                        NS      l6.nstld.com.
                        NS      a6.nstld.com.
                        NS      c6.nstld.com.
                        NS      d6.nstld.com.
                        NS      f6.nstld.com.
                        NS      m6.nstld.com.
(
			NS	m7.nstld.com.
)
$ORIGIN name.
0-0onlus                NS      ns7.ehiweb.it.
                        NS      ns8.ehiweb.it.
0-g                     MX      10 mx01.nic
                        MX      10 mx02.nic
                        MX      10 mx03.nic
                        MX      10 mx04.nic
$ORIGIN 0-g.name
moutamassey             NS      ns01.yahoodomains.jp.
                        NS      ns02.yahoodomains.jp.
`
	for x := range dns.ParseZone(strings.NewReader(zone), "", "") {
		if x.Error != nil {
		   fmt.Println(x.Error)
		} else {
		  // Do something with x.RR
			fmt.Println(x.RR)
			fmt.Println(x.RR.Header())

			fmt.Printf("Name:%s, Type:%s,Class:%s,RDLen:%d,TTL:%d\n ",x.Header().Name, dns.TypeToString[x.Header().Rrtype],dns.ClassToString[x.Header().Class], x.Header().Rdlength,x.Header().Ttl)
	      }
	}
}
