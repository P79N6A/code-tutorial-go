package main

import "fmt"

func main() {
        fmt.Println(findMin([]int{5,1,2,3,4}))
        fmt.Println(findMin([]int{5,1}))
}

func findMin(nums []int) int {
        if len(nums)==1{return nums[0]}
        if nums[0]<nums[len(nums)-1]{return nums[0]} // 递增
        l,r := 0,len(nums)
        for l<r {
                m := (l+r)/2
                // 想想为什么会有两个出口?
                if m>0&&nums[m]<nums[m-1]{ // 找到最小值
                        return nums[m]
                }
                if m<len(nums)-1&& nums[m]>nums[m+1] { // 找到最大值
                        return nums[m+1]
                }
                if nums[m]>nums[0] { //二分
                        l=m+1
                } else {
                        r=m-1
                }
        }
        return nums[r]
}