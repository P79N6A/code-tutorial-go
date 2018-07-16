package main

import "fmt"

func lowerBound(nums []int,target int) int {
    l,r := 0,len(nums)
    for ;l<r; {
        m := (l+r)/2
        //和upperbound唯一的区别
        // < 意味着,l会右移到小于target,那么r就是第一个等于target的下标,如果没有等于就是大于
        if nums[m] < target {
            l = m + 1
        } else {
            r = m
        }
    }
    return r
}
func upperBound(nums []int,target int) int {
    l,r := 0,len(nums)
    for ;l<r; {
        m := (l+r)/2
        // <= 意味着,即使等于target,l也会继续右移,那么r就是第一个大于target的下标
        if nums[m] <= target {
            l = m + 1
        } else {
            r = m
        }
    }
    return r
}
func binarySearch(nums []int,target int) int {
    l,r := 0,len(nums)  // 初始值，注意r
    for ;l<r; {
        // 有可能越界
        m := l + (l-r)/2
        if nums[m] == target {
            return m
        }else if nums[m] <= target {
            l = m + 1
        } else {
            r = m
        }
    }
    return -1
}




func lower_bound(arr []int,x int) int {
    l,r := 0,len(arr)
    for l < r {
        m := (l+r)/2
        if arr[m] < x {
            l = m + 1
        } else {
            r = m
        }
    }
    return r
}

func main() {
    fmt.Println(upperBound([]int{1,2,3,3,3,3,4,5},5))
    a := []int{1,2,3,4,5}
    copy(a[1:],a[2:])
    fmt.Println(a)
    
}
