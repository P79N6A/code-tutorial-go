package main

import (
        "github.com/kayon/qqwry"
        "log"
)
func main() {
        qw := qqwry.New("/Users/gaolichuang/workspace/go/src/github.com/kayon/qqwry/qqwry.dat")
        res := qw.Search("123.12.3.1")
        log.Printf("\nIP: %s\nBegin: %s\nEnd: %s\n%s %s\n", res.IP, res.Begin, res.End, res.Country, res.Area)
}
