package main

import (
    "fmt"
    "strconv"
)

/*
Description
Given a sorted integer array where the range of elements are in the inclusive range [lower, upper], return its missing ranges.


Example
Given nums = [0, 1, 3, 50, 75], lower = 0 and upper = 99
return ["2", "4->49", "51->74", "76->99"].
*/
func findMissingRanges(nums []int, low,up int) []string {
    ret := make([]string,0)
    l := low
    r := 0
    for i:=0;i<=len(nums);i++ {
        if i < len(nums) && nums[i] <= up {
            r = nums[i]
        } else {
            r = up
        }
        if l==r {
            l += 1
        } else if r > l {
            if r - l == 1 {
                ret = append(ret,strconv.Itoa(l))
            } else {
                if r == up {
                    ret = append(ret,strconv.Itoa(l) + "->" + strconv.Itoa(r))
                } else {
                    ret = append(ret,strconv.Itoa(l) + "->" + strconv.Itoa(r-1))
                }
            }
            l = r + 1
        }
    }
    return ret
}
func main() {
    fmt.Println(findMissingRanges([]int{1,2,3,4},2,3))
    fmt.Println(findMissingRanges([]int{1,2,4},2,5))
    fmt.Println(findMissingRanges([]int{1,2,10,20},2,20))
}
