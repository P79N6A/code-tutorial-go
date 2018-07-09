package main

import (
        "gcodebase/http_lib"
        "time"
        "fmt"
        "gcodebase/time_util"
)

func getipurl(ims,num int,url string, c chan int) {
        suc := 0
        for i := 0; i < num;i++ {
                body, err := http_lib.GetUrl("GET", url, "")
                fmt.Printf("%v:%s\n",err,body)
                if err == nil {
                        suc += 1
                }
                time.Sleep(time.Millisecond * time.Duration(ims))
        }
        c <- suc
}
func main() {
        t1 := time_util.GetTimeInMs()
        mc := make([]chan int,0)
        for i := 0;i < 50; i++ {
                c := make(chan int)
                mc = append(mc,c)
                go getipurl(1,100,"http://127.0.0.1:9800/IpInfo/211.93.83.107",c)
        }
        all := 0
        for _,c := range mc {
                all += <- c
        }

        fmt.Printf("%d use %d ms",all,time_util.GetTimeInMs() - t1)

}
