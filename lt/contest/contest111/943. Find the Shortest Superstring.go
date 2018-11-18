package main

import (
    "math"
    "strings"
    "fmt"
)

func main() {
    fmt.Println(shortestSuperstring([]string{"catg","ctaagt","gcta","ttca","atgcatc"}))
}

func shortestSuperstring(A []string) string {
    seq := make([]int,len(A))
    for i:=0;i<len(A);i++ {seq[i]=i}
    ans := math.MaxInt64
    anss := ""
    xx := []int{}
    for _,ss := range permute(seq) {
        str := A[ss[0]]
        for i:=1;i<len(ss);i++ {
            j:= 0
            for j=0;j<len(A[ss[i]]);j++ {
                if !strings.HasSuffix(str, A[ss[i]][:j]) {
                    break
                }
            }
            str += A[ss[i]][(j-1):]
        }
        if len(str) < ans {
            ans = len(str)
            anss = str
            xx = ss
        }
    }
    fmt.Println(xx)
    return anss
}

func permute(nums []int) [][]int {
    result := make([][]int,0)
    pem(nums,0,&result)
    return result
}
func pem(nums []int,start int, result *[][]int) {
    if start == len(nums)-1 {
        n := make([]int,len(nums))
        copy(n,nums)
        *result = append(*result,n)
        return
    }
    for i:=start;i<len(nums);i++ {
        nums[start],nums[i]=nums[i],nums[start]
        // 和组合的最大区别，第二个参数传入的是start+1
        pem(nums,start+1,result)
        nums[start],nums[i]=nums[i],nums[start]
    }
}