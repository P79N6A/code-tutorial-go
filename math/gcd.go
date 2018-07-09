package main

import "fmt"

func GCDSlice(s []int32) int32 {
        if len(s) <= 0 {
                return 1
        }
        gcd := s[0]
        for i:=1;i < len(s);i++ {
                gcd = CalGCD(gcd,s[i])
        }
        return gcd
}
func CalGCD(a, b int32) int32 {
        for b > 0 {
                a, b = b, a%b
        }
        return a
}
func main() {
        fmt.Println(GCDSlice([]int32{64,8,16,32}))
        fmt.Println(GCDSlice([]int32{}))
}
