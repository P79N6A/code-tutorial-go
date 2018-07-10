package main

import (
        "bytes"
        "io"
        "os"
        "fmt"
)

func main() {
        buf := bytes.NewBuffer(nil)
        f, _ := os.Open("/Users/zhujunchao/workspace/go/sync-d.sh") // Error handling elided for brevity.
        io.Copy(buf, f)           // Error handling elided for brevity.
        f.Close()
        s := string(buf.Bytes())
        fmt.Println(s)
}
