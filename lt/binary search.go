package main

import "fmt"

func upperbound(a []int, t int) int {
    l,r := 0,len(a)
    for ;l<r;{
        m := (l + r)/2
        if a[m] <= t {
            l = m + 1
        } else {
            r = m
        }
    }
    return r
}
func lowerbound(a []int, t int) int {
    l,r := 0,len(a)
    for ;l<r;{
        m := (l + r)/2
        if a[m] < t {
            l = m + 1
        } else {
            r = m
        }
    }
    return r
}
func binarysearch(a []int, t int) int {
    left,right := 0,len(a)
    for {
        if left >= right {
            break
        }
        mid := left + (right - left) / 2
        if a[mid] < t {
            left = mid + 1
        } else {
            right = mid
        }
        fmt.Println(right,left,mid)
    }
    return right
}

func main() {
    fmt.Println(upperbound([]int{1,2,4,6,7,7,8,9},7))
    fmt.Println(lowerbound([]int{1,2,4,6,7,7,8,9},7))

}
