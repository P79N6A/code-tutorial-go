package main

import (
    "strconv"
    "fmt"
)
/*
Given a sorted integer array where the range of elements are [0, 99] inclusive, return its missing ranges.
For example, given [0, 1, 3, 50, 75], return [“2”, “4->49”, “51->74”, “76->99”]
*/
func findMissingRanges(num []int,low,high int) []string {
    ret := make([]string,0)
    l,r := low,0
    for i:=0;i<len(num);i++ {
        r = num[i]
        if r > high {
            r = high
            i = len(num)
        }
        if l == r {
            l += 1
        } else if r > l {
            if r-l == 1 {
                ret = append(ret,strconv.Itoa(l))
            } else {
                ret = append(ret,strconv.Itoa(l) + "->" + strconv.Itoa(r-1))
            }
            l = r+1
        }
    }
    if len(num)>0 {
        l = num[len(num)-1] + 1
        r = high
        if r > l {
            if r-l == 1 {
                ret = append(ret,strconv.Itoa(l))
            } else {
                ret = append(ret,strconv.Itoa(l) + "->" + strconv.Itoa(r))
            }
        }
    }
    return ret
}
func main() {
    fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75},0,99))
    fmt.Println(findMissingRanges([]int{1, 3, 50, 75},0,99))
    fmt.Println(findMissingRanges([]int{1, 3, 50, 75},0,74))
    fmt.Println(findMissingRanges([]int{0, 1, 3, 50, 75,99},0,99))
    fmt.Println(findMissingRanges([]int{1, 3, 50, 75},0,3))
}