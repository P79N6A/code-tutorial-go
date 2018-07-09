package main

import (
        "fmt"
        "time"
        "strings"
)

func main() {
        var vv []byte
        fmt.Println("aaaa",string(vv))
        str := "xxx86xxx"
        idx := strings.Index(str,"86")
        fmt.Println(idx)
        fmt.Println(str[:idx])
        a := "smart/nxnop5004.json"
        fmt.Println(a[6:len(a)-4])
        y,m,d := time.Now().Date()
        fmt.Println(y)
        fmt.Println(int(m))
        fmt.Println(d)
        fmt.Printf("%.2d",11)

        fmt.Println(strings.HasSuffix("XXXXX",""))
        fmt.Println(strings.TrimLeft("aabcds","*"))
        fmt.Println(strings.Replace(a,"xn","--",-1))
        fmt.Println(strings.ToLower("Xi'an"))

        a = Repeat("xxx",10)
        fmt.Println(a,len(a))
}

func Repeat(s string, count int) string {
        // Since we cannot return an error on overflow,
        // we should panic if the repeat will generate
        // an overflow.
        // See Issue golang.org/issue/16237
        b := make([]byte, len(s)*count)
        // 长度有限制，如果超过了，返回较小的。
        bp := copy(b, s)
        for bp < len(b) {
                copy(b[bp:], b[:bp])
                bp *= 2
                fmt.Println(bp,string(b),s)
        }
        return string(b)
}