package main

import "fmt"

/*
Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.

Example 1:
Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
Output: True
Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.
Note:

1 <= k <= len(nums) <= 16.
0 < nums[i] < 10000.
*/

func main() {
    fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1},4))//true
    fmt.Println(canPartitionKSubsets([]int{5,2,5,5,5,5,5,5,5,5,5,5,5,5,5,3},15))//true
    fmt.Println(canPartitionKSubsets([]int{2,2,2,2,3,4,5},4))// false
    fmt.Println(canPartitionKSubsets([]int{129,17,74,57,1421,99,92,285,1276,218,1588,215,369,117,153,22}, 3))
                                          //[1 1  1  1   1   1  0   0   0   1    1    1   0   1   1  0]
    fmt.Println(canPartitionKSubsets([]int{57,1421,285,1588,215,369,153}, 2))
}
func canPartitionKSubsets(nums []int, k int) bool {
    sum := 0
    for _,n := range nums {sum += n}
    if sum%k != 0 {return false} // 是否能够整除
    return canParti(nums,0,0,0,0,k,sum/k,make(map[string]bool))
}
// 因为nums长度给了16，所以visit可以使用int的bit来做，否则只能使用[]int了
func canParti(nums []int,visit int,start,cursum int,curnum int,k,sum int,memo map[string]bool) bool {
    key := fmt.Sprintf("%d-%d",visit,cursum) // memo key
    if _,ok := memo[key];ok {
        return memo[key]
    }
    if k == 1 {return true}
    if cursum == sum && curnum > 0 { // curnum是为了处理如果sum=0的时候仍有有效num选出来
        return canParti(nums,visit,0,0,0,k-1,sum,memo)
    }
    for i:=start;i<len(nums);i++ {
        if visit>>uint(i) & 1 == 0 {
            if cursum + nums[i] <= sum {
                // start or start+1 带来的性能差距较大，后者可以优化接近20ms！！！！
                if canParti(nums, visit|1<<uint(i), start+1,cursum+nums[i], curnum+1, k, sum,memo) {
                    memo[key]=true
                    return true
                }
            }
        }
    }
    memo[key]=false
    return false
}
