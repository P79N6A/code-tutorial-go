package main

import (
        "sort"
        "fmt"
)

func main() {
        fmt.Println(sumSubseqWidths([]int{2,1,3}))
        // 136988321
        fmt.Println(sumSubseqWidths([]int{96,87,191,197,40,101,108,35,169,50,168,182,95,80,144,43,18,60,174,13,77,173,38,46,80,117,13,19,11,6,13,118,39,80,171,36,86,156,165,190,53,49,160,192,57,42,97,35,124,200,84,70,145,180,54,141,159,42,66,66,25,95,24,136,140,159,71,131,5,140,115,76,151,137,63,47,69,164,60,172,153,183,6,70,40,168,133,45,116,188,20,52,70,156,44,27,124,59,42,172}))
}

func sumSubseqWidths(A []int) int {
        sort.Ints(A)
        cache := make(map[int]uint64)
        cache[0]=1
        for i:=1;i<20001;i++ {
                cache[i] = (cache[i-1]*2)% 1000000007
        }
        var ret uint64
        for i:=0;i<len(A);i++ {
                for j:=i+1;j<len(A);j++ {
                        aa := uint64(A[j] - A[i])
                        aa = (aa * cache[j-i-1])% 1000000007
                        ret = (ret + aa) % 1000000007
                }
        }
        return int(ret % 1000000007)
}
func sumSubseqWidths1(A []int) int {
        sort.Ints(A)
        var ret uint64
        for i:=0;i<len(A);i++ {
                for j:=i+1;j<len(A);j++ {
                        aa := uint64(A[j]-A[i])
                        for ii:=0;ii<j-i-1;ii++ {
                                aa = (aa*2) % 1000000007
                        }
                        ret = (ret + aa)%1000000007
                }
        }
        return int(ret % 1000000007)
}