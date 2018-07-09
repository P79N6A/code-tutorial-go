package main
import (
        "bytes"
        "time"
        "os"
        "io"
        "github.com/miekg/dns"
        "fmt"
)
func ReadFileToString(name string) ([]byte, error) {
        buf := bytes.NewBuffer(nil)
        f, err := os.Open(name)
        if err != nil {
                return []byte(""), err
        }
        defer f.Close()
        _,err = io.Copy(buf, f)
        if err != nil {
                return []byte(""), err
        }
        return buf.Bytes(),nil
}
func NowT() int64 {
        return time.Now().UnixNano() / int64(1000000)
}
func main() {
        t1 := NowT()
        content,_ := ReadFileToString("/Users/zhujunchao/workspace/go/src/glory_dns/mdata/etc/glory_dns/zone/h0.com.hosts")
        for x := range dns.ParseZone(bytes.NewReader(content), "", "") {
                fmt.Println(x.Header().Name)
        }
        fmt.Printf("use %d ms\n",NowT() - t1)
}
