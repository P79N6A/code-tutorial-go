package main

import (
        "strings"
        "fmt"
        "path/filepath"
)

func main() {
        fmt.Println(len(strings.Split(",",",")))
        fmt.Println((strings.Split("",",")))
        newf := filepath.Join("http://api.dn-dns.gls.acadn.com:9116/idns/config/v2/platforms/nxnop061","xxx")
        fmt.Println(newf)
        fmt.Println(strings.Contains("aaa","aaa"))


        ret := strings.Split("a   b"," ")
        fmt.Println(len(ret))


        ss := strings.TrimRight("dns1.flxdns.com.","flxdns.com.")
        ss = strings.Replace("dns1.flxdns.com.","flxdns.com.","",-1)
        s := "a.dwf.ewf.dns1.flxdns.com."
        sub := "flxdns.com."
        ssss := s[:len(s) - len(sub)]
        fmt.Println(ssss)
        fmt.Println(ss)
        fmt.Println(sub[:0])
}
