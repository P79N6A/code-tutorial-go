package main

import (
        "fmt"
        "strings"
        "time"
        "github.com/davecgh/go-spew/spew"
)

type OBJ struct {
        L []string  `protobuf:"bytes,4,rep,name=ips" json:"ips,omitempty"`
        M       string  `protobuf:"bytes,1,opt,name=mm" json:"mm,omitempty"`
        I int  `protobuf:"bytes,1,opt,name=ii" json:"ii,omitempty"`
}
func GetSuperiorDomain(domain string) string {
        if domain == "." {
                return ""
        }
        dm := strings.Split(domain, ".")
        if len(dm) == 2 {
                return "."
        }
        return strings.Join(dm[1:], ".")
}
func main() {
        fmt.Println(GetSuperiorDomain("a.com."))
        fmt.Println(GetSuperiorDomain("x.x.x.a.com."))
        fmt.Println(GetSuperiorDomain("ab.a.com."))
        fmt.Println(GetSuperiorDomain("com."))
        fmt.Println(GetSuperiorDomain("."))
        obj := OBJ{
                L:[]string{"22","333","xxx"},
                M:"efwoij",
                I:12,
        }
        fmt.Println(fmt.Sprintf("%s",spew.Sdump(obj)))
        fmt.Println(time.Now().Format("15:04:05"))
        fmt.Println(cap([1024*1024]byte{}))
}
