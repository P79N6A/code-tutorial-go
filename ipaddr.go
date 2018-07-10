package main

import (
        "strings"
        "net"
        "fmt"
)



func GetParentDomain(domain string) string {
        if domain == "." {
                return ""
        }
        idx := 0
        for idx < len(domain) && domain[idx] != '.' {
                idx += 1
        }
        // the last dot or not
        if idx < len(domain)-1 {
                idx += 1
        }
        return domain[idx:]
}
func IP2Uint(ip net.IP) uint64 {
        nip := ip.To4()
        return uint64(nip[0])<<24 | uint64(nip[1])<<16 |
        uint64(nip[2])<<8 | uint64(nip[3])
}
// 1.2.3.4 = > 00000001 00000010 00000011 00000100
func IP2BinarySeries(ip net.IP) [32]byte {
        ipi := IP2Uint(ip)
        binIP := [32]byte{0}
        idx := 31
        for ipi != 0 {
                binIP[idx] = byte(ipi & 1)
                ipi = ipi >> 1
                idx -= 1
        }
        return binIP
}
func BinarySeries2IP(ipi [32]byte) net.IP {
        i := 0
        a,b,c,d := byte(0),byte(0),byte(0),byte(0)
        for {
                a += ipi[i]
                i += 1
                if i == 8 {
                        break
                }
                a = a << 1
        }
        for {
                b += ipi[i]
                i += 1
                if i == 16 {
                        break
                }
                b = b << 1
        }
        for {
                c += ipi[i]
                i += 1
                if i == 24 {
                        break
                }
                c = c << 1
        }
        for {
                d += ipi[i]
                i += 1
                if i == 32 {
                        break
                }
                d = d << 1
        }
        return net.IPv4(a,b,c,d)
}

func GenCIDRFromIP(ip1,ip2 string) string {
        ip1Ser := IP2BinarySeries(net.ParseIP(ip1))
        ip2Ser := IP2BinarySeries(net.ParseIP(ip2))
        fmt.Println(ip1Ser)
        fmt.Println(ip2Ser)
        i := 0
        for ;i < len(ip1Ser);i++ {
                if ip1Ser[i] != ip2Ser[i] {
                        break
                }
        }
        mask := i
        for ;i < 32;i++ {
                ip1Ser[i] = byte(0)
        }
        fmt.Println(ip1Ser)
        ipi := BinarySeries2IP(ip1Ser)
        return fmt.Sprintf("%s/%d",ipi.String(),mask)
}
func CIDR2BinarySeries(cidr string) (error, []byte){
        ip, ipnet, err := net.ParseCIDR(cidr)
        if err != nil {
                return err, []byte{0}
        }
        ipi := IP2Uint(ip)
        mask, _ := ipnet.Mask.Size()
        binCIDR := [32]byte{0}
        idx := 31
        for ipi != 0 {
                binCIDR[idx] = byte(ipi & 1)
                ipi = ipi >> 1
                idx -= 1
        }

        return nil, binCIDR[:mask]
}
func test() {
        ip := net.IPv4(byte(1),byte(255),byte(128),byte(4))
        fmt.Println(ip.String())
        fmt.Println(IP2BinarySeries(ip))
        fmt.Println(BinarySeries2IP(IP2BinarySeries(ip)).String())
        fmt.Println(GenCIDRFromIP("1.2.3.4","1.3.0.0"))
        _,ret := CIDR2BinarySeries("1.1.0.0/8")
        fmt.Println(ret)
}
const kViewSeparator = '-'
func GetLevelByViewName(view string) int {
        level := 1
        for _,s := range view {
                if s == kViewSeparator {
                        level += 1
                }
        }
        return level
}
func GetUpLevelViewName(view string) string {
        idx := strings.LastIndex(view,"-")
	if idx < 0 {

	}
        return view[:idx]
}
func main() {
        fmt.Println(GetLevelByViewName("cnc-bj"))
        fmt.Println(GenCIDRFromIP("223.204.228.0","223.204.28.255"))
        fmt.Println(GenCIDRFromIP("255.255.255.255","255.255.255.255"))
        fmt.Println(GetUpLevelViewName("cnc-bj-few-fwe"))
        fmt.Println(GetUpLevelViewName("cnc"))
}