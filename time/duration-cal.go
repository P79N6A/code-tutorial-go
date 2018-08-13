package main

import (
        "time"
        "fmt"
    "strconv"
)

func main() {
        var a time.Duration = 100
        a = a * 98 / 100
        fmt.Println(a)

        d,e := time.ParseDuration("11h")
        fmt.Println(d,e)

    parserDuration := func (str string) (error,time.Duration) {
        unitMap := map[byte]time.Duration {
            'w':time.Hour*24*7,
            'd':time.Hour*24,
        }
        var ret time.Duration
        pos := 0
        for i:=0;i<len(str);i++ {
            if str[i]>='0' && str[i]<='9' {
                continue
            }
            if _,ok := unitMap[str[i]];!ok {
                return fmt.Errorf("invalid unit %s",string(str[i])),0
            }
            n,e := strconv.Atoi(str[pos:i])
            if e != nil {
                return fmt.Errorf("parse int error %s %v",str[pos:i],e),0
            }
            ret += time.Duration(n) * unitMap[str[i]]
            pos = i+1
        }
        return nil,ret
    }
    fmt.Println(parserDuration("1w"))
    fmt.Println(parserDuration("1w2d"))
    fmt.Println(parserDuration("1d2w"))
    fmt.Println(parserDuration("1wd"))
    fmt.Println(parserDuration("1ww"))
    fmt.Println(parserDuration(""))
}
