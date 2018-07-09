package main

import (
        "gcodebase/http_lib"
        "fmt"
        "net/url"
)

func main() {
        u,_ := url.Parse("HTTPS://127.0.0.1:9113/status")
        fmt.Println(u.Scheme)
        body,err := http_lib.GetUrl("GET","https://127.0.0.1:9113/statusi","")
        fmt.Println(body)
        fmt.Println(err)
}
