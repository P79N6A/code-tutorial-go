package main

import (
    "fmt"
    "git.byted.org/content/cloud/scheduler/utils/time_util"
    "time"
    "net"
)

func main() {
    fmt.Println((-1+3)%3)
    a := 123
    f := 1.23456
    fmt.Printf("%+10d\n", a) //+123
    fmt.Printf("%.10d\n", a) //+123
    fmt.Printf("%.4f\n", f)  //+000000123，利用０来补齐位数，而不是空格
    fmt.Printf("%02d\n", 1)
    fmt.Printf("%02d\n", 1)
    fmt.Printf("%v", time.Second)
    _, _, e := net.SplitHostPort("127.0.0.1")
    //fmt.Println(h)
    //fmt.Println(p)
    fmt.Println(e)
    var sss []string
    fmt.Println(len(sss))
    fmt.Println(^(99 - 1))
    fmt.Println(99 ^ 0)
    fmt.Println(99 ^ 1)
    aa := "123456789"
    fmt.Println(aa[:5])

    fmt.Printf( "%[2]v %[1]s\n","xx","yy")

    fmt.Printf("%[3]*.[2]*[1]f\n", 12.0, 2, 6)

    fmt.Println("Xxxxxxxx")
    for i:=0;i<24;i++ {
        fmt.Printf("%02d\n",i)
    }
    var aaa uint64 = 999
    fmt.Printf("%014x///%6x",aaa,time_util.GetCurrentTimeStamp())

}
