package heap

import "fmt"
/*
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/
func findKthLargest(nums []int, k int) int {
        l,r := 0,len(nums)-1
        for l <= r {
                p := partition(nums,l,r)
                fmt.Println(l,r,k,p)
                if p == k - 1 {
                        // 必须有找到的情况,因为找到和没找到返回内容不一样,nums[p],nums[r]
                        return nums[p]
                } else if p < k - 1 {
                        l = p + 1
                } else {
                        r = p-1
                }
        }
        return nums[r]
}
func partition(nums []int, start,end int) int {
        piviot,l,r := nums[start],start+1,end
        // 边界比较多,
        for l <= r {
                if nums[l] < piviot && nums[r] > piviot {
                        nums[l], nums[r] = nums[r], nums[l]
                        l += 1
                        r -= 1
                }
                fmt.Println(l,r,piviot,start,end)
                if nums[l] >= piviot {
                        l += 1
                }
                if nums[r] <= piviot {
                        r -= 1
                }
        }
        nums[start],nums[r] = nums[r],nums[start]
        return r
}

func main() {
        fmt.Println(findKthLargest([]int{3,2,1,5,6,4},2))
        //fmt.Println(findKthLargest([]int{2,1},1))
}
