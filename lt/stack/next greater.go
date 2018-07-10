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
func nextGreaterElements2(nums []int) []int {
    /*
    从右往左递减栈；解决右侧第一个比元素大的值
    处理在入栈过程中，stack不需要保持下标
    */
    ret := make([]int, len(nums))
    for i := 0; i < len(ret); i++ {
        ret[i] = -1
    }
    stack := make([]int, 0)
    for i := 2*len(nums) - 1; i >= 0; i-- {
        cur := nums[i%len(nums)]
        for len(stack) > 0 && cur >= stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) <= 0 || cur < stack[len(stack)-1] {
            if len(stack) > 0 {
                ret[i%len(nums)] = stack[len(stack)-1]
            }
            // stack存的值
            stack = append(stack, cur)
        }
    }
    return ret
}
func nextGreaterElements3(nums []int) []int {
    /*
    从左到右的递减栈，解决右侧第一个比元素大的问题
    处理在出栈过程中，stack需要存储下标，因为入栈后在比较大小会用到
    */
    ret := make([]int, len(nums))
    stack := make([]int, 0)
    for i := 0; i < 2*len(nums); i++ {
        cur := nums[i%len(nums)]
        for len(stack) > 0 && cur > nums[stack[len(stack)-1]] {
            ret[stack[len(stack)-1]] = cur
            stack = stack[:len(stack)-1]
        }
        if len(stack) <= 0 || cur <= nums[stack[len(stack)-1]] {
            stack = append(stack, i%len(nums))
        }
    }
    if len(stack) > 0 {
        // 处理-1 的情况
        max := nums[stack[0]]
        for i := 0; i < len(stack); i++ {
            if max == nums[stack[i]] {
                ret[stack[i]] = -1
            } else {
                break
            }
        }
    }
    return ret
}
func main() {
fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
