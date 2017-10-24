package main

import "fmt"

type LogS struct {
        L int
}
func (l *LogS)Debug(s string) {
        if l.L > 3 {
                fmt.Print(s)
        }
}
var logs *LogS
func VLog(i int) *LogS {
        logs.L = i
        return logs
}
type AAA struct {

}
func (a *AAA)String() string {
        fmt.Println("XXXX")
        return "A"
}
func main() {
        logs = new(LogS)
        a := new(AAA)
        VLog(2).Debug(a.String())
}
