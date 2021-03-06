package main

import "fmt"

/*
Suppose you have N integers from 1 to N.
We define a beautiful arrangement as an array
that is constructed by these N numbers successfully
if one of the following is true for the ith position (1 <= i <= N) in this array:

The number at the ith position is divisible by i.
i is divisible by the number at the ith position.
Now given N, how many beautiful arrangements can you construct?

Example 1:
Input: 2
Output: 2
Explanation:

The first beautiful arrangement is [1, 2]:

Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).

Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).

The second beautiful arrangement is [2, 1]:

Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).

Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.
Note:
N is a positive integer and will not exceed 15.
*/
func countArrangement(N int) int {
    nums := make([]int,0)
    for i:=0;i<=N;i++ {
        nums = append(nums,i)
    }
    n := 0
    dfsC(nums,1,&n)
    return n
}
/*
生成排列，生成过程中剪枝
*/
func dfsC(nums []int,start int,num *int) {
    if start == len(nums) {
        fmt.Println(nums)
        *num += 1
        return
    }
    for i:=start;i<len(nums);i++ {
        if nums[i]%start == 0 || start%nums[i]==0 {
            nums[i], nums[start] = nums[start], nums[i]
            dfsC(nums, start+1, num)
            nums[i], nums[start] = nums[start], nums[i]
        }
    }
}
func main() {
    fmt.Println(countArrangement(3))
    
}
