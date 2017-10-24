package main

import (
        "crypto/sha256"
        "fmt"
        "io"
)

func main() {
        h := sha256.New()
        io.WriteString(h, "fewfewfew")
        fmt.Printf("%x", h.Sum(nil))
}
