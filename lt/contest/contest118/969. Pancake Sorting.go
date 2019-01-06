package main

import "fmt"

/*
Given an array A, we can perform a pancake flip:
We choose some positive integer k <= A.length,
then reverse the order of the first k elements of A.
We want to perform zero or more pancake flips (doing them one after another in succession) to sort the array A.

Return the k-values corresponding to a sequence of pancake flips that sort A.
Any valid answer that sorts the array within 10 * A.length flips will be judged as correct.



Example 1:

Input: [3,2,4,1]
Output: [4,2,4,3]
Explanation:
We perform 4 pancake flips, with k values 4, 2, 4, and 3.
Starting state: A = [3, 2, 4, 1]
After 1st flip (k=4): A = [1, 4, 2, 3]
After 2nd flip (k=2): A = [4, 1, 2, 3]
After 3rd flip (k=4): A = [3, 2, 1, 4]
After 4th flip (k=3): A = [1, 2, 3, 4], which is sorted.
Example 2:

Input: [1,2,3]
Output: []
Explanation: The input is already sorted, so there is no need to flip anything.
Note that other answers, such as [3, 3], would also be accepted.


Note:

1 <= A.length <= 100
A[i] is a permutation of [1, 2, ..., A.length]
*/

func main() {
    fmt.Println(pancakeSort([]int{3,2,4,1}))
    fmt.Println(pancakeSort([]int{2,1,3}))
    fmt.Println(pancakeSort([]int{3,1,2}))
}
/*
答案应该有很多种，这种方式相对好理解，就是每次把最大的一个换到末尾
方法 :
0.初始化最大值为Len(A)
1.找到本轮最大下标
2.1 如果最大在末尾，不处理，最大下标-1，重复
2.2 如果最大在第一个，整体翻转，将最大的放到最后，重复
2.3 最大在中间，先将最大的翻转到第一个；在整体翻转，重复
*/

func pancakeSort(A []int) []int {
    ans := make([]int,0)
    for n:=len(A);!inorder(A,n);n--{
        mi := maxi(A,n)
        if mi < n-1 {
            if mi > 0 {
                ans = append(ans,mi+1)
                reverse(&A,0,mi)
            }
            ans = append(ans, n)
            reverse(&A, 0, n-1)
        }
    }
    return ans
}
func inorder(A []int,n int) bool {
    for i:=1;i<n;i++ {
        if A[i-1] > A[i] {
            return false
        }
    }
    return true
}
func maxi(A []int,n int) int {
    m,mi := A[0],0
    for i:=1;i<n;i++ {
        if A[i] > m {
            m = A[i]
            mi = i
        }
    }
    return mi
}
func reverse(A *[]int,s,e int) {
    for s < e {
        (*A)[s],(*A)[e]=(*A)[e],(*A)[s]
        s += 1
        e -= 1
    }
}
