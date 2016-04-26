package main
import (
    "fmt"
    "net"
    "sync"
    "time"
    . "github.com/miekg/dns"
)


func HelloServer(w ResponseWriter, req *Msg) {
    m := new(Msg)
    m.SetReply(req)

    m.Extra = make([]RR, 1)
    m.Extra[0] = &TXT{Hdr: RR_Header{Name: m.Question[0].Name, Rrtype: TypeTXT, Class: ClassINET, Ttl: 0}, Txt: []string{"Hello world"}}
    w.WriteMsg(m)
}

func AnotherHelloServer(w ResponseWriter, req *Msg) {
    m := new(Msg)
    m.SetReply(req)

    m.Extra = make([]RR, 1)
    m.Extra[0] = &TXT{Hdr: RR_Header{Name: m.Question[0].Name, Rrtype: TypeTXT, Class: ClassINET, Ttl: 0}, Txt: []string{"Hello example"}}
    w.WriteMsg(m)
}
func RunLocalUDPServerWithFinChan(laddr string) (*Server, string, chan struct{}, error) {
    pc, err := net.ListenPacket("udp", laddr)
    if err != nil {
        return nil, "", nil, err
    }
    server := &Server{PacketConn: pc, ReadTimeout: time.Hour, WriteTimeout: time.Hour}

    waitLock := sync.Mutex{}
    waitLock.Lock()
    server.NotifyStartedFunc = waitLock.Unlock

    fin := make(chan struct{}, 0)

    go func() {
        server.ActivateAndServe()
        close(fin)
        pc.Close()
    }()

    waitLock.Lock()
    return server, pc.LocalAddr().String(), fin, nil
}
func RunLocalUDPServer(laddr string) (*Server, string, error) {
    server, l, _, err := RunLocalUDPServerWithFinChan(laddr)

    return server, l, err
}

func main() {
    HandleFunc("miek.nl.", HelloServer)
    HandleFunc("example.com.", AnotherHelloServer)
    //defer HandleRemove("miek.nl.")
    //defer HandleRemove("example.com.")

    s, addrstr, err := RunLocalUDPServer("127.0.0.1:53")
    if err != nil {
        fmt.Printf("unable to run test server: %v", err)
    }
    fmt.Printf("Address server %s",addrstr)
    defer s.Shutdown()
    for true {
        time.Sleep(time.Second)
    }
}
