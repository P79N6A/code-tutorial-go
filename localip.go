package main

import (
        "net"
        "fmt"
)
func GetLocalIP(name string) net.IP {
        ifaces, err := net.Interfaces()
        if err != nil {
                return nil
        }
        for _, iface := range ifaces {
                addrs,err := iface.Addrs()
                if err != nil {
                        continue
                }
                if name != "" && iface.Name != name {
                        continue
                }
                for _,addr := range addrs {
                        //fmt.Println(addr)
                        // check the address type and if it is not a loopback the display it
                        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
                                if ipnet.IP.To4() != nil {
                                        return ipnet.IP
                                }
                        }
                }
        }
        return nil
}
func main() {
        fmt.Println(GetLocalIP(""))
        /*
        ipn := net.ParseIP("10.100.100.10")
        ipn = ipn.To4()
        fmt.Println([]byte(ipn)[0] == 10)
        ipa := net.IP{}
        fmt.Println(ipa.To4())
        */

}
/*
func main() {
        ifaces, err := net.Interfaces()
        // handle err
        if err != nil {
                fmt.Println(err)
        }
        for _, i := range ifaces {
                fmt.Println(i.Name)

                addrs, err := i.Addrs()
                // handle err
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println(addrs)
                for _, addr := range addrs {
                        var ip net.IP
                        switch v := addr.(type) {
                        case *net.IPNet:
                                ip = v.IP
                        case *net.IPAddr:
                                ip = v.IP
                        }
                        // process IP address
                        fmt.Println(ip)

                }

        }
}
*/
