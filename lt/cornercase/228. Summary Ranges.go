package cornercase

import "fmt"

/*
Given a sorted integer array without duplicates, return the summary of its ranges.

Example 1:

Input:  [0,1,2,4,5,7]
Output: ["0->2","4->5","7"]
Explanation: 0,1,2 form a continuous range; 4,5 form a continuous range.
Example 2:

Input:  [0,2,3,4,6,8,9]
Output: ["0","2->4","6","8->9"]
Explanation: 2,3,4 form a continuous range; 8,9 form a continuous range.
*/

func summaryRanges(nums []int) []string {
    ret := make([]string,0)
    p,q := 0,0
    for q < len(nums) {
        for q+1 < len(nums) {
            if nums[q+1] == nums[q] + 1 {
                q += 1
            } else {
                break
            }
        }
        if p == q {
            ret = append(ret,fmt.Sprintf("%d",nums[p]))
        } else {
            ret = append(ret,fmt.Sprintf("%d->%d",nums[p],nums[q]))
        }
        p=q+1
        q=p
    }
    return ret
}

func main() {
    fmt.Println(summaryRanges([]int{0,1,2,4,5,7}))
    fmt.Println(summaryRanges([]int{0,2,3,4,6,8,9}))
    fmt.Println(summaryRanges([]int{0,1,2,3}))
    fmt.Println(summaryRanges([]int{}))

    
}
