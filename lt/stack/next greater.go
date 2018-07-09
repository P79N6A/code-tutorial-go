package main

import "fmt"

func nextGreaterElements(nums []int) []int {
        ret := make([]int,len(nums))
        for i:=0;i<len(ret);i++ {ret[i]=-1}
        stack := make([]int,0)
        for i:=2*len(nums)-1;i>=0;i-- {
                cur := nums[i%len(nums)]
                for len(stack) > 0 && cur >= stack[len(stack)-1] {
                        stack = stack[:len(stack)-1]
                }
                if len(stack)<=0 || cur < stack[len(stack)-1] {
                        if len(stack) > 0 {
                                ret[i%len(nums)]=stack[len(stack)-1]
                        }
                        stack = append(stack,cur)
                }
        }
        return ret
}
func nextGreaterElements3(nums []int) []int {
        ret := make([]int,len(nums))
        stack := make([]int,0)
        for i:=0;i<2*len(nums);i++ {
                cur := nums[i%len(nums)]
                //fmt.Println(stack,cur,i,ret)
                for len(stack) > 0 && cur > nums[stack[len(stack)-1]] {
                        ret[stack[len(stack)-1]] = cur
                        stack = stack[:len(stack)-1]
                }
                if len(stack)<= 0 || cur <= nums[stack[len(stack)-1]] {
                        stack = append(stack,i%len(nums))
                }
        }
        if len(stack)>0 {
                max := nums[stack[0]]
                for i:=0;i<len(stack);i++ {
                        if max == nums[stack[i]] {
                                ret[stack[i]] = -1
                        } else {
                                break
                        }
                }
        }

        return ret
}
func nextGreaterElements2(nums []int) []int {
        ret := make([]int,len(nums))
        stack := make([]int,0)
        for i:=0;i<2*len(nums);i++ {
                cur := nums[i%len(nums)]
                //fmt.Println(stack,cur,i,ret)
                for len(stack) > 0 && cur > nums[stack[len(stack)-1]] {
                        ret[stack[len(stack)-1]] = cur
                        stack = stack[:len(stack)-1]
                }
                if len(stack)<= 0 || cur <= nums[stack[len(stack)-1]] {
                        stack = append(stack,i%len(nums))
                }
        }
        for len(stack) > 0 {
                ret[stack[len(stack)-1]] = -1
                stack = stack[:len(stack)-1]
        }
        return ret
}
func main() {
        fmt.Println(nextGreaterElements([]int{1,2,1}))
}
