package main

import "fmt"

/*
A zero-indexed array A of length N contains all integers from 0 to N-1. Find and return the longest length of set S, where S[i] = {A[i], A[A[i]], A[A[A[i]]], ... } subjected to the rule below.

Suppose the first element in S starts with the selection of element A[i] of index = i, the next element in S should be A[A[i]], and then A[A[A[i]]]… By that analogy, we stop adding right before a duplicate element occurs in S.

Example 1:
Input: A = [5,4,0,3,1,6,2]
Output: 4
Explanation:
A[0] = 5, A[1] = 4, A[2] = 0, A[3] = 3, A[4] = 1, A[5] = 6, A[6] = 2.

One of the longest S[K]:
S[0] = {A[0], A[5], A[6], A[2]} = {5, 6, 2, 0}
Note:
N is an integer within the range [1, 20,000].
The elements of A are all distinct.
Each element of A is an integer within the range [0, N-1].
*/

func main() {
    fmt.Println(arrayNesting([]int{5,4,0,3,1,6,2}))
    fmt.Println(arrayNesting([]int{0,1,2,3,4}))

}
/*
问存在的最长的nest array长度多少。
有点union find 的意思。
由于是0-n 直接用nums[i]*-1 对0的处理就不够好；采用赋值方式标记
*/
func arrayNesting(nums []int) int {
    /*
    有两种边界条件
    1.自己跟自己成环:  1->1  算1个
    2.多人参与  1->2->3->1 算3个
    这种好像do while 更合适一些，用for写，j==nums[j]的case少加了一次1
    */
    visit := len(nums)
    ans := 0
    for i:=0;i<len(nums);i++ {
        if nums[i]==visit {continue}
        one := 1
        for j:=i;j<visit&&nums[j]<visit&&j!=nums[j]; {
            one += 1
            if j==nums[j] { // 需要+1后在退出
                break
            }
            nj := nums[j]
            nums[j] = visit
            j = nj
        }
        ans = max(ans,one)
    }
    return ans
}
func max(a,b int) int {
    if a>b {return a}
    return b
}