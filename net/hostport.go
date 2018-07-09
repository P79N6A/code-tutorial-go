package main

import (
        "strings"
        "strconv"
        "fmt"
        "net"
)

func ParseListenAddr(addr string, defaultPort int) (error,string) {
        if strings.Contains(addr,":") {
                h, p, err := net.SplitHostPort(addr)
                if err != nil {
                        return err,""
                }
                pi, e := strconv.Atoi(p)
                if e != nil {
                        return err,""
                }
                return nil,fmt.Sprintf("%s:%d",h,pi)
        } else {
                if nil == net.ParseIP(addr) {
                        return fmt.Errorf("%s not vaild addr",addr),""
                }
                return nil,fmt.Sprintf("%s:%d",addr,defaultPort)
        }

}
func main() {
        err,h := ParseListenAddr("1.2.3",88)
        fmt.Println(err)
        fmt.Println(h)
}
