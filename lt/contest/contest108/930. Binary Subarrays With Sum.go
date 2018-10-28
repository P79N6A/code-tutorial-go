package main

import "fmt"

func main() {
    fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1},2))
    fmt.Println(numSubarraysWithSum([]int{0,0,0,0,0},0))
}
func numSubarraysWithSum(A []int, S int) int {
    ans := 0
    sd := make(map[int]int)
    sd[0]=1
    sum := 0
    for _,a := range A {
        sum += a
        if _,ok := sd[sum-S];ok {
            ans += sd[sum-S]
        }
        sd[sum]+=1
    }
    return ans
}