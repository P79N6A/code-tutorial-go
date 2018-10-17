package main

import (
    "net/http"
    "net/http/httptrace"
    "fmt"
)

func main() {

    req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
    trace := &httptrace.ClientTrace{
        GotConn: func(connInfo httptrace.GotConnInfo) {
            fmt.Printf("Got Conn: %+v\n", connInfo)
        },
        DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
            fmt.Printf("DNS Info: %+v\n", dnsInfo)
        },
    }


    req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))


    client := new(http.Client)
    _, err := client.Do(req)


    //_, err := http.DefaultTransport.RoundTrip(req)
    if err != nil {
        fmt.Println(err)
    }
}
