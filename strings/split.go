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
}
