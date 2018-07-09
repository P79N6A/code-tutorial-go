package main

import (
        "gcodebase/file"
        "strings"
        "fmt"
)

func main() {
        filename := "/Users/zhujunchao/Downloads/gnop2502.2__61.163.252.26.log"
        ans := make(map[string]int)
        file.FileLineReader(filename,"#",func(line string) {
                ret := strings.Split(line," ")
                rett := strings.Split(ret[1],":")
                fmt.Printf("%s:%s  %s\n",rett[0],rett[1],ret[13])
                key := fmt.Sprintf("%s:%s",rett[0],rett[1])
                if ret[13] != "NULL" {
                        ans[key] += 1

                }
        })
        for k,v := range ans {
                fmt.Printf("%s\t%d\n",k,v)
        }
}
