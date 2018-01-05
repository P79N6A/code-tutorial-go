package main

import (
        "fmt"
        "time"
        "strings"
)

func main() {
        var vv []byte
        fmt.Println(string(vv))
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
}
