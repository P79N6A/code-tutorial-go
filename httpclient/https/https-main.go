package main

import (
        // "fmt"
        // "io"
        "net/http"
        "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("This is an example server.\n"))
        // fmt.Fprintf(w, "This is an example server.\n")
        // io.WriteString(w, "This is an example server.\n")
}

func main() {
        http.HandleFunc("/hello", HelloServer)
        err := http.ListenAndServeTLS(":1234", "server1.pem", "server1.key", nil)
        if err != nil {
                log.Fatal("ListenAndServe: ", err)
        }
}