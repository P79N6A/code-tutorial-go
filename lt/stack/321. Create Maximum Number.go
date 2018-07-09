package main

import "fmt"

/*

Given two arrays of length m and n with digits 0-9 representing two numbers. Create the maximum number of length k <= m + n from digits of the two. The relative order of the digits from the same array must be preserved. Return an array of the k digits.

Note: You should try to optimize your time and space complexity.

Example 1:

Input:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
Output:
[9, 8, 6, 5, 3]
Example 2:

Input:
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
Output:
[6, 7, 6, 0, 4]
Example 3:

Input:
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
Output:
[9, 8, 9]
*/
func maxNumInArryC(nums []int,k int) []int {
    drop := len(nums)-k
    ret := make([]int,0)
    for i:=0;i<len(nums);i++ {
        for drop > 0 && len(ret) > 0 && nums[i]>ret[len(ret)-1] {
            ret = ret[:len(ret)-1]
            drop -= 1
        }
        ret = append(ret,nums[i])
    }
    for drop > 0 && len(ret) > 0 {
        ret = ret[:len(ret)-1]
        drop -= 1
    }
    return ret
}
func maxNumInArryC1(nums []int,k int) []int {
    if k == 0 {return nil}
    if k >= len(nums){return nums}
    max := 0
    res :=make([]int,0)
    ret :=make([]int,0)
    maxNumInArry(nums,0,1,k,&max,&res,&ret)
    return ret
}
func maxNumInArry(nums []int,pos,sum int,k int,max *int, res *[]int,ret *[]int) {
    if k == 0 && sum > *max {
        *max = sum
        *ret = make([]int,0)
        for i:=0;i<len(*res);i++ {
            *ret = append(*ret,(*res)[i])
        }
        return
    }
    for i:=pos;i<len(nums);i++ {
        *res = append(*res,nums[i])
        maxNumInArry(nums,i+1,10*sum+nums[i],k-1,max,res,ret)
        *res = (*res)[:len(*res)-1]
    }
}
func isbiggerdifflen(nums1,nums2 []int)bool {
    if len(nums1) <= 0 {return false}
    if len(nums2) <= 0 {return true}
    if nums1[0] > nums2[0] {
        return true
    } else if nums1[0]<nums2[0] {
        return false
    }
    return isbiggerdifflen(nums1[1:],nums2[1:])
}
func mergeNum(nums1,nums2 []int) []int {
    ret := make([]int,len(nums1)+len(nums2))
    i,j,idx := 0,0,0
    n1,n2 := nums1,nums2
    for len(n1) + len(n2)>0{
        if isbiggerdifflen(n1,n2) {
            ret[idx] = n1[0]
            i+=1
            idx+=1
            n1 = n1[1:]
        } else {
            ret[idx] = n2[0]
            j+=1
            idx+=1
            n2 = n2[1:]
        }
    }
    return ret
}
func maxNumber(nums1 []int, nums2 []int, k int) []int {
    m,n := len(nums1),len(nums2)
    start := 0
    if k-n > 0 {
        start = k-n
    }
    end := k
    if m < k {
        end=m
    }
    fmt.Println(start,end)
    var max []int
    for i:=start;i<=end;i++ {
        fmt.Println(i,k-i)
        x := mergeNum(maxNumInArryC(nums1,i),maxNumInArryC(nums2,k-i))
        fmt.Println(x)
        if max == nil || isbigger(x,max) {
            max=x
        }
    }
    return max
}
func isbigger(x,max []int) bool {
    for i:=0;i<len(x);i++ {
        if x[i] > max[i] {
            return true
        } else if x[i] < max[i] {
            return false
        }
    }
    return true
}


func main() {
    //fmt.Println(maxNumber([]int{3, 4, 6, 5},[]int{9, 1, 2, 5, 8, 3},5))
    //fmt.Println(maxNumber([]int{6,7},[]int{6,0,4},5))
    //[2,5,6,4,4,0]
    //[7,3,8,0,6,5,7,6,2]
    //15
    //fmt.Println(maxNumber([]int{2,5,6,4,4,0},[]int{7,3,8,0,6,5,7,6,2},15))
    //[6,7,5]
    //[4,8,1]
    fmt.Println(maxNumber([]int{6,7,5},[]int{4,8,1},3))
    //fmt.Println(maxNumInArryC([]int{1,2,3},0))
}
