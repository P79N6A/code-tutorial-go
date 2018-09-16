package main
import (
        "gcodebase/hash"
        "path/filepath"
    "strings"
    "fmt"
)
func GenWriteFileTemporaryName(name string) string {
        dir := filepath.Dir(name)
        base := filepath.Base(name)
        subfix := hash.GenerateKey(5)
        newBase := "." + base + "." + subfix
        return filepath.Join(dir,newBase)
}
func main() {
//        str := GenWriteFileTemporaryName("./../xxx.log")
//        fmt.Println(str)
    fmt.Println(HigherDomain("."))
    fmt.Println(HigherDomain("com."))
    fmt.Println(HigherDomain("www.dnion.com."))
    fmt.Println(HigherDomain(""))
    fmt.Println(HigherDomain("xxx"))
}

func HigherDomain(domain string) string {
    idx := strings.Index(domain,".")
    if idx <= 0 {return "."}
    if idx >= len(domain)-1 {
        return "."
    }
    return domain[(idx+1):]
}