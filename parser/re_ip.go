package main

import (
        "regexp"
        "fmt"
)

var msg = `
;; opcode: QUERY, status: NOERROR, id: 25202
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 5, ADDITIONAL: 6

;; QUESTION SECTION:
;netcenteropt.xdwscache.ourwebcdn.com.	IN	 A

;; ANSWER SECTION:
netcenteropt.xdwscache.ourwebcdn.com.	60	IN	A	222.132.5.100

;; AUTHORITY SECTION:
ourwebcdn.com.	86400	IN	NS	dns3.ourwebcdn.org.
ourwebcdn.com.	86400	IN	NS	dns1.ourwebcdn.org.
ourwebcdn.com.	86400	IN	NS	dns2.ourwebcdn.info.
ourwebcdn.com.	86400	IN	NS	dns4.ourwebcdn.info.
ourwebcdn.com.	86400	IN	NS	dns5.ourwebcdn.org.

;; ADDITIONAL SECTION:
dns3.ourwebcdn.org.	7200	IN	A	211.94.114.110
dns1.ourwebcdn.org.	7200	IN	A	61.130.24.30
dns2.ourwebcdn.info.	7200	IN	A	122.138.54.196
dns4.ourwebcdn.info.	7200	IN	A	222.163.201.19
dns5.ourwebcdn.org.	7200	IN	A	60.222.223.78

;; OPT PSEUDOSECTION:
; EDNS: version 0; flags: ; udp: 4096
; SUBNET: 121.192.44.0/32/19

;; query time: 33985 µs, server: dns1.ourwebcdn.org.:53(udp), size: 296 bytes [Compress]
`
func main() {
        numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
        regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock
        r, _ := regexp.Compile(regexPattern)
        for _,v := range r.FindAllString(msg,-1) {
                fmt.Println(v)
        }
        newMsg := r.ReplaceAllStringFunc(msg, func(info string) string{
                return "xxx"
        })
        fmt.Println(newMsg)

}

