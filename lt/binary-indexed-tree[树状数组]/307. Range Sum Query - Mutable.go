package main

import "fmt"

/*
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.

Example:

Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
Note:

The array is only modifiable by the update function.
You may assume the number of calls to update and sumRange function is distributed evenly.
*/

type NumArray struct {
    num []int
    // binary indexed tree
    bit []int
}


func Constructor(nums []int) *NumArray {
    na := &NumArray{
        num:make([]int,len(nums)+1),
        bit:make([]int,len(nums)+1),
    }
    for i:=0;i<len(nums);i++ {
        na.Update(i,nums[i])
    }
    return na
}


func (this *NumArray) Update(i int, val int)  {
    diff := val - this.num[i+1]
    for j:=i+1;j<len(this.num);j+=j&-j {
        this.bit[j] += diff
    }
    this.num[i+1]=val
}

func (this *NumArray) getSum(i int) int {
    ret := 0
    for j:=i;j>0;j-=j&-j {
        ret += this.bit[j]
    }
    return ret
}

func (this *NumArray) SumRange(i int, j int) int {
    return this.getSum(j+1)-this.getSum(i)
}

func main() {
    obj := Constructor([]int{1,3,5});
    fmt.Println(obj.SumRange(0,2))
    obj.Update(1,2)
    fmt.Println(obj.SumRange(0,2))
    fmt.Println(^2)
    fmt.Println(^3)
    fmt.Println(^(4)+1)
}
