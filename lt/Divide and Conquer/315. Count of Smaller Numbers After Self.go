package main

import (
    "fmt"
    "sort"
)

/*
You are given an integer array nums and you have to return a new counts array.
The counts array has the property where counts[i] is the number of smaller elements
to the right of nums[i].

Example:

Input: [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.
*/
func countSmaller(nums []int) []int {
    // 位置会变化，所以需要有一个只读的数组
    rnums := make([]int,len(nums))
    copy(rnums,nums)
    result := make([]int,len(nums))
    mergesort(nums,rnums,0,len(nums),&result)
    return result
}
func mergesort(nums,rnums []int, start,end int,ret *[]int) {
    if end - start <= 1 {
        return
    }
    mid := (start+end)/2
    mergesort(nums,rnums,start,mid,ret)
    mergesort(nums,rnums,mid,end,ret)
    for i:=start;i<mid;i++ {
        j:=mid
        for ;j<end;j+=1 {
            // binary search
            if nums[j]>=rnums[i] {
                break
            }
        }
        (*ret)[i]+= j-mid
        /*
        for j:=mid;j<end;j++ {
            if nums[j] < rnums[i] {(*ret)[i]+=1}
        }
        //*/
    }
    sort.Ints(nums[start:end])
}
func main() {
    fmt.Println(countSmaller([]int{26,78,27,100,33,67,90,23,66,5,38,7,35,23,52,22,83,51,98,69,81,32,78,28,94,13,2,97,3,76,99,51,9,21,84,66,65,36,100,41}))
    //[10,27,10,35,12,22,28,8,19,2,12,2,9,6,12,5,17,9,19,12,14,6,12,5,12,3,0,10,0,7,8,4,0,0,4,3,2,0,1,0]
}
