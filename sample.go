package main

import (
        "fmt"
        "strings"
)

func main() {
        a := "xxx"
        b := "xxxx-www"
        aa := strings.Split(a,"-")
        fmt.Println(len(aa))
        fmt.Println(aa)
        bb := strings.Split(b,"-")
        fmt.Println(bb)

        CIDRs := make([]string,0)
        fmt.Printf("%s\t{%s;%s;};","XXX", "XXXXX",strings.Join(CIDRs,";"))
}
