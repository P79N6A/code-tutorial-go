package heap

import (
        "math"
        "fmt"
)
/*
You are given two integer arrays nums1 and nums2 sorted in ascending order and an integer k.

Define a pair (u,v) which consists of one element from the first array and one element from the second array.

Find the k pairs (u1,v1),(u2,v2) ...(uk,vk) with the smallest sums.

Example 1:
Given nums1 = [1,7,11], nums2 = [2,4,6],  k = 3

Return: [1,2],[1,4],[1,6]

The first 3 pairs are returned from the sequence:
[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
Example 2:
Given nums1 = [1,1,2], nums2 = [1,2,3],  k = 2

Return: [1,1],[1,1]

The first 2 pairs are returned from the sequence:
[1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
Example 3:
Given nums1 = [1,2], nums2 = [3],  k = 3

Return: [1,3],[2,3]

All possible pairs are returned from the sequence:
[1,3],[2,3]
Credits:
Special thanks to @elmirap and @StefanPochmann for adding this problem and creating all test cases.

https://leetcode.com/problems/merge-k-sorted-lists/description/
跟这个题目类似
*/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
        ret := make([][]int,0)
        idx := make([]int,len(nums1))
        for {
                left := false
                min,mini,minj := math.MaxInt64,0,0
                for i:=len(nums1)-1;i>=0;i-- {
                        if idx[i] < len(nums2) && (nums1[i]+nums2[idx[i]]) <= min {
                                min = nums1[i]+nums2[idx[i]]
                                mini,minj,left = i,idx[i],true
                        }
                }
                if left {
                        ret = append(ret,[]int{nums1[mini],nums2[minj]})
                        if len(ret) >= k {
                                return ret
                        }
                        idx[mini] += 1

                } else {
                        break
                }
        }
        return ret
}

func main() {
        //fmt.Println(kSmallestPairs([]int{1,7,11},[]int{2,4,6},2))
        //fmt.Println(kSmallestPairs([]int{1,7,11},[]int{},1))
//        fmt.Println(kSmallestPairs([]int{1,2,3,4,5},[]int{3,5,7,9},20))
        fmt.Println(kSmallestPairs([]int{3,22,35,56,76},[]int{3,22,35,56,76},25))
}
