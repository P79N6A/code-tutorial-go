package main

import "fmt"
/*
Given a n x n matrix where each of the rows and columns are sorted in ascending order,
find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
Note:
You may assume k is always valid, 1 ≤ k ≤ n2.
*/
func kthSmallest(matrix [][]int, k int) int {
        row := len(matrix)
        if row <= 0 {return 0}
        col := len(matrix[0])
        // 找到l,r
        // l,r的范围就是矩阵中的最大值和最小值，由于行列有序，也就是第一个和最后一个元素
        l,r := matrix[0][0],matrix[row-1][col-1]
        for l < r {
                mid := (l + r)/2
                /* 子函数部分，
                因为找的是排名第K个，所以找大于mid有多少个，
                进而确认下排名第K个在l-mid,还是mid-r，实现二分
                 */
                counter := 0
                for i:=0;i<row;i++ {
                        // upper bound 第一个大于mid的;并不是等于,
                        // 因为等于可能无效,所以找第一个比他大的肯定有效
                        counter += upperBound(matrix[i],mid)
                }
                fmt.Println(l,r,mid,counter)
                //没有等于的场景,因为必须将序列走完.等于的可能是没有结果的.
                // 不可能返回mid的,因为mid未必是有效的
                //lower bound,找第一个不小于k的[大于等于的]
                if counter < k {
                        l = mid + 1
                } else {
                        r = mid
                }
        }
        //结果答案必须是在r或者的,其他的并不一定有效
        return r
}
func upperBound(nums []int,target int) int {
        l,r := 0,len(nums)
        for l < r {
                mid := l + (r-l)/2
                if nums[mid] <= target {
                        l = mid + 1
                } else {
                        r = mid
                }
        }
        return r
}

func main() {
        matrix := [][]int {
                []int{1,5,9},
                []int{10,11,13},
                []int{12,13,15},
        }
        fmt.Println(kthSmallest(matrix,8))
        fmt.Println(binsearch([]int{1,2,2,2,2,3,4},2))
}
