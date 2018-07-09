package main

import (
        "fmt"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,
                "Hi, This is an example of https service in golang!")
}

func main() {
        http.HandleFunc("/", handler)
        err := http.ListenAndServeTLS(":8081", "server1.pem",
                "server1.key", nil)
        if err != nil {
                fmt.Println(err)
        }
}