package main

import "fmt"

func genPalindrome1(n int) []int {
    //  是需要给结果组合的，需要判断是否包含其他长度的。
    if n == 0 {return nil}
    if n == 1 {return []int{0,1,2,3,4,5,6,7,8,9}}
    if n == 2 {return []int{0,11,22,33,44,55,66,77,88,99}}
    ret := make([]int,0)
    sub := genPalindrome(n-2,true)
    for i:=0;i<=9;i++ {
        for _,s := range sub{
            r := i
            for j:=1;j<n;j++ {
                r *=10
            }
            rt := r+s*10+i
            ret = append(ret, rt)
        }
    }
    return ret
}

func genPalindrome(n int,isSub bool) []int {
    //  是需要给结果组合的，需要判断是否包含其他长度的。
    if n == 0 {return nil}
    if n == 1 {return []int{0,1,2,3,4,5,6,7,8,9}}
    if n == 2 {return []int{0,11,22,33,44,55,66,77,88,99}}
    ret := make([]int,0)
    sub := genPalindrome(n-2,true)
    //ret = append(ret,0)
    if isSub {
        for _,s := range sub{
            r := 0
            for j:=1;j<n;j++ {
                r *=10
            }
            rt := r+s*10+0
            ret = append(ret, rt)
        }
    }
    for i:=1;i<=9;i++ {
        for _,s := range sub{
            r := i
            for j:=1;j<n;j++ {
                r *=10
            }
            rt := r+s*10+i
            ret = append(ret, rt)
        }
    }
    return ret
}

func main() {
    fmt.Println(genPalindrome(5,false))
}
