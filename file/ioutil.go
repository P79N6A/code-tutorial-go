package main

import (
        "fmt"
        "os"
)

func main() {
        path := "/Users/gaolichuang/workspace/go/src/glory_dns/logs"
        file,e := os.Open(path)
        fmt.Println(e)
        defer file.Close()
        fs,e := file.Readdir(-1)
        fmt.Println(e)
        for _,f := range fs {
                fmt.Println(f.Name())
                fmt.Println(f.IsDir())
        }
        fmt.Println(file.Name())
}
